package auth

const Model = `
	[request_definition]
	r = action, role
	
	[policy_definition]
	p = action, role
	
	[role_definition]
	g = _, _
	
	[policy_effect]
	e = some(where (p.eft == allow))
	
	[matchers]
	m = g(r.action, p.action) && r.role == p.role
`
