package main

import "slices"

// accountsMerge takes a list of accounts (each with a name and emails),
// and merges accounts that share at least one common email.
// Returns a list of merged accounts with sorted emails.
func accountsMerge(accounts [][]string) [][]string {
	// Number of accounts
	n := len(accounts)

	// Initialize a Disjoint Set (Union-Find) to group connected accounts
	ds := NewDS(n)

	// Map each email to the index of the first account it was seen in
	// This helps us detect when an email appears in multiple accounts
	emailMaps := map[string]int{}

	// First pass: iterate over each account and its emails
	for i, account := range accounts {
		emails := account[1:] // All elements after the name are emails

		for _, email := range emails {
			// Check if this email has been seen before
			parent, exist := emailMaps[email]

			if !exist {
				// First time seeing this email: map it to current account index
				emailMaps[email] = i
			} else {
				// Email already seen in another account: merge the two accounts
				// Union current account (i) with the account that previously had this email (parent)
				ds.UnionBySize(i, parent)
			}
		}
	}

	// After unions, group all emails by their ultimate parent account
	mergedEmails := make([][]string, n)

	// Second pass: for each email, find the root (representative) account
	// and collect all emails under that root
	for email, node := range emailMaps {
		parent := ds.FindUltimateParent(node)

		// Initialize the slice if not already done
		if len(mergedEmails[parent]) == 0 {
			mergedEmails[parent] = []string{}
		}

		// Append the email to the merged list of the root account
		mergedEmails[parent] = append(mergedEmails[parent], email)
	}

	// Final result container
	ans := [][]string{}

	// Third pass: build the final merged account list
	for parent, emails := range mergedEmails {
		// Skip empty groups (no emails merged here)
		if len(emails) == 0 {
			continue
		}

		// Sort emails lexicographically as required
		slices.Sort(emails)

		// Start with the account name from the original account at index 'parent'
		account := []string{accounts[parent][0]}
		// Append all sorted emails
		account = append(account, emails...)

		// Append a new slice to ans
		ans = append(ans, account)
	}

	return ans
}

// DisjointSet implements Union-Find with both size and rank (though only size is used here)
type DisjointSet struct {
	parents []int // parent pointers for each node
	sizes   []int // size of component rooted at each node (used in UnionBySize)
	ranks   []int // rank for union by rank (unused in this solution)
}

// NewDS creates a new DisjointSet with n nodes (one per account)
func NewDS(n int) *DisjointSet {
	parents := make([]int, n)
	sizes := make([]int, n)
	ranks := make([]int, n)

	// Initially, each node is its own parent (self-loop), size 1, rank 0
	for i := range n {
		parents[i] = i
		sizes[i] = 1
		ranks[i] = 0
	}

	return &DisjointSet{
		parents: parents,
		sizes:   sizes,
		ranks:   ranks,
	}
}

// FindUltimateParent finds the root of the set containing 'node'
// Uses path compression: flattens the tree during lookup for efficiency
func (d *DisjointSet) FindUltimateParent(node int) int {
	if node == d.parents[node] {
		return node // root found
	}

	// Recursively find and update parent to root (path compression)
	d.parents[node] = d.FindUltimateParent(d.parents[node])
	return d.parents[node]
}

// UnionByRank merges two sets by comparing ranks (not used in accountsMerge)
// Lower rank tree attaches to higher rank tree. If equal, one is chosen and rank increases.
func (d *DisjointSet) UnionByRank(u, v int) {
	pu := d.FindUltimateParent(u)
	pv := d.FindUltimateParent(v)

	if pu == pv {
		return // already in same set
	}

	// Attach smaller rank tree under larger rank tree
	if d.ranks[pu] < d.ranks[pv] {
		d.parents[pu] = pv
	} else if d.ranks[pv] < d.ranks[pu] {
		d.parents[pv] = pu
	} else {
		// Equal ranks: attach pu under pv, increase pv's rank
		d.parents[pu] = pv
		d.ranks[pv]++
	}
}

// UnionBySize merges two sets by size: smaller set attaches to larger set
// Updates size of the new root accordingly
// This is the method used in accountsMerge for efficient union
func (d *DisjointSet) UnionBySize(u, v int) {
	pu := d.FindUltimateParent(u)
	pv := d.FindUltimateParent(v)

	if pu == pv {
		return // already connected
	}

	// Attach smaller component to larger one
	if d.sizes[pu] < d.sizes[pv] {
		d.parents[pu] = pv         // pu's root now points to pv
		d.sizes[pv] += d.sizes[pu] // pv's size increases
	} else {
		d.parents[pv] = pu         // pv's root points to pu
		d.sizes[pu] += d.sizes[pv] // pu's size increases
	}
}
