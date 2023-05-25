package ocigatewayexample

import (
	cfg "github.com/mheers/ocigateway"
)

version: "2.0"
log: {}
storage: ""

http: {
  addr: ":5000"
}

cacheDir: "/tmp/demo/cache"

logLevel: "debug"
logFormat: "json"

cfg.#ociGateway & {
	gateways: [
		cfg.#ftpGateway & {
			name:     "my-ftp"
			host:     "ftp.myhost.com"
			port:     21
			username: "Marcel"
			password: "1234"
		},
		cfg.#ftpGateway & {
			name:     "my-ftps"
			username: "Marcel"
			password: "1234"
			host:     "ftp.myhost.com"
			port:     21
		},
		cfg.#gitGateway & {
			name:     "my-git"
			repoURL:  "my-repo"
			// sshKey:   "my-ssh-key"
			// sshKeyFile:   "my-ssh-key"
			password: "1234"
		},
		cfg.#sftpGateway & {
			name:     "my-sftp"
			host:     "sftp.myhost.com"
			port:     22
			username: "Marcel"
			password: "1234"
			insecure: true // if false we need to set hostKey
		},
	]
}
