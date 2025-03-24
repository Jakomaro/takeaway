Key Directives
Directive	Example	Purpose
max-age	max-age=3600	Cache is fresh for 1 hour (in seconds)
no-cache	no-cache	Cache must revalidate with server before reuse (no stale responses)
no-store	no-store	Disallow caching entirely (for sensitive data)
public	public, max-age=3600	Allow caching by shared caches (e.g., CDNs)
private	private, max-age=3600	Allow caching only in user's browser (not intermediate systems)
must-revalidate	must-revalidate, max-age=3600	Force revalidation after max-age expires