// Package cxdomainutils provides utilities for domain matching,
// including support for wildcard domains.
package cxdomainutils

import (
    "strings"
)

// MatchDomain checks if the domain matches the given pattern.
// The pattern can be a specific domain or a wildcard pattern
// such as "*.example.com". If the pattern starts with "*.",
// it matches any subdomain of the specified domain.
func MatchDomain(pattern, domain string) bool {
    if strings.HasPrefix(pattern, "*.") {
        return strings.HasSuffix(domain, pattern[1:])
    }
    return domain == pattern
}
