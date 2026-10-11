package cluster

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseInfo preserves all INFO fields, including replication offset diagnostics.
func ParseInfo(raw string) map[string]string {
	out := make(map[string]string)
	for _, line := range strings.Split(raw, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if key, value, ok := strings.Cut(line, ":"); ok {
			out[key] = value
		}
	}
	return out
}

// ParseNodes parses the administrative CLUSTER NODES wire format. Importing and
// migrating entries are retained separately: they do not assign slot ownership.
func ParseNodes(raw string) ([]NodeState, error) {
	var nodes []NodeState
	seen := make(map[string]bool)
	for i, line := range strings.Split(strings.TrimSpace(raw), "\n") {
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		if len(fields) < 8 {
			return nil, fmt.Errorf("CLUSTER NODES line %d: expected at least 8 fields", i+1)
		}
		if seen[fields[0]] {
			return nil, fmt.Errorf("CLUSTER NODES line %d: duplicate node ID %s", i+1, fields[0])
		}
		seen[fields[0]] = true
		n := NodeState{ID: fields[0], Addr: endpoint(fields[1]), Flags: strings.Split(fields[2], ","), LinkState: fields[7]}
		if fields[3] != "-" {
			n.PrimaryID = fields[3]
		}
		var err error
		if n.PingSent, err = strconv.ParseInt(fields[4], 10, 64); err != nil || n.PingSent < 0 {
			return nil, fmt.Errorf("CLUSTER NODES line %d: invalid ping timestamp", i+1)
		}
		if n.PongReceived, err = strconv.ParseInt(fields[5], 10, 64); err != nil || n.PongReceived < 0 {
			return nil, fmt.Errorf("CLUSTER NODES line %d: invalid pong timestamp", i+1)
		}
		if n.ConfigEpoch, err = strconv.ParseUint(fields[6], 10, 64); err != nil {
			return nil, fmt.Errorf("CLUSTER NODES line %d: invalid epoch", i+1)
		}
		if n.Primary() && n.Replica() {
			return nil, fmt.Errorf("CLUSTER NODES line %d: conflicting role flags", i+1)
		}
		if n.LinkState != "connected" && n.LinkState != "disconnected" {
			return nil, fmt.Errorf("CLUSTER NODES line %d: invalid link state", i+1)
		}
		for _, token := range fields[8:] {
			if strings.HasPrefix(token, "[") {
				if err := validateTransition(token); err != nil {
					return nil, fmt.Errorf("CLUSTER NODES line %d: %w", i+1, err)
				}
				n.SlotTransitions = append(n.SlotTransitions, token)
				continue
			}
			slots, err := parseRange(token)
			if err != nil {
				return nil, fmt.Errorf("CLUSTER NODES line %d: %w", i+1, err)
			}
			n.Slots = append(n.Slots, slots)
		}
		nodes = append(nodes, n)
	}
	if len(nodes) == 0 {
		return nil, fmt.Errorf("CLUSTER NODES returned no nodes")
	}
	return nodes, nil
}

func endpoint(raw string) string {
	addr, _, _ := strings.Cut(raw, "@")
	addr, _, _ = strings.Cut(addr, ",")
	return addr
}

func parseRange(token string) (SlotRange, error) {
	lo, hi, isRange := strings.Cut(token, "-")
	start, err := strconv.Atoi(lo)
	if err != nil {
		return SlotRange{}, fmt.Errorf("invalid slot %q", token)
	}
	end := start
	if isRange {
		end, err = strconv.Atoi(hi)
		if err != nil {
			return SlotRange{}, fmt.Errorf("invalid slot range %q", token)
		}
	}
	if start < 0 || end >= SlotCount || end < start {
		return SlotRange{}, fmt.Errorf("out-of-range slot %q", token)
	}
	return SlotRange{Start: start, End: end}, nil
}

func validateTransition(token string) error {
	if !strings.HasSuffix(token, "]") {
		return fmt.Errorf("invalid slot transition %q", token)
	}
	body := token[1 : len(token)-1]
	for _, separator := range []string{"->-", "-<-"} {
		if slot, id, ok := strings.Cut(body, separator); ok {
			r, err := parseRange(slot)
			if err != nil || r.Start != r.End || id == "" {
				return fmt.Errorf("invalid slot transition %q", token)
			}
			return nil
		}
	}
	return fmt.Errorf("invalid slot transition %q", token)
}

// AnalyzeCoverage detects gaps and overlapping claims, even when a failed
// primary still advertises slots. Duplicate claims by one node are also invalid.
func AnalyzeCoverage(nodes []NodeState) Coverage {
	var c Coverage
	owners := make([]string, SlotCount)
	claims := make([]int, SlotCount)
	unhealthy := make([]bool, SlotCount)
	for _, n := range nodes {
		for _, r := range n.Slots {
			// Guard public callers constructing NodeState without ParseNodes.
			if r.Start < 0 || r.End >= SlotCount || r.Start > r.End {
				c.InvalidClaims = append(c.InvalidClaims, fmt.Sprintf("%s invalid slot range %d-%d", n.ID, r.Start, r.End))
				continue
			}
			for slot := r.Start; slot <= r.End; slot++ {
				claims[slot]++
				if owners[slot] == "" {
					owners[slot] = n.ID
				}
				if !n.Primary() || !n.Healthy() {
					unhealthy[slot] = true
				}
			}
		}
	}
	for _, count := range claims {
		if count > 0 {
			c.Assigned++
		}
	}
	c.Missing = rangesMatching(func(slot int) bool { return claims[slot] == 0 })
	c.Duplicate = rangesMatching(func(slot int) bool { return claims[slot] > 1 })
	c.UnhealthyOwners = rangesMatching(func(slot int) bool { return unhealthy[slot] })
	for slot := 0; slot < SlotCount; {
		if claims[slot] != 1 {
			slot++
			continue
		}
		start, owner := slot, owners[slot]
		for slot+1 < SlotCount && claims[slot+1] == 1 && owners[slot+1] == owner {
			slot++
		}
		c.Owners = append(c.Owners, OwnerRange{Start: start, End: slot, RedisID: owner})
		slot++
	}
	return c
}

func rangesMatching(match func(int) bool) []SlotRange {
	var ranges []SlotRange
	for slot := 0; slot < SlotCount; slot++ {
		if !match(slot) {
			continue
		}
		start := slot
		for slot+1 < SlotCount && match(slot+1) {
			slot++
		}
		ranges = append(ranges, SlotRange{Start: start, End: slot})
	}
	return ranges
}

func ownerMap(nodes []NodeState) []string {
	owners := make([]string, SlotCount)
	for _, r := range AnalyzeCoverage(nodes).Owners {
		for slot := r.Start; slot <= r.End; slot++ {
			owners[slot] = r.RedisID
		}
	}
	return owners
}
