package auth

const CommonPolicy Policy = `
	p, /api/VerifyToken:POST, admin, shipper
	p, /api/CreateCategory:POST, admin, shipper
`
