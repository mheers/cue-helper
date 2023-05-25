package ocigateway

import (
	"list"
)

#ociGateway: {
	...
	gateways: [...#gateways]

	_checks: {
		#names: [ for n in gateways {n.name}]
		uniqueNames: [ for i, x in #names if !list.Contains(list.Drop(#names, i+1), x) {x}]
		isUniqe: len(uniqueNames) == len(#names)
		isUniqe: true
	}
}

ociGateway: #ociGateway & {}

#gateway: {
	name: string
	type: "file" | "http" | "s3" | "docker" | "git" | "ftp" | "nfs" | "sftp" | "smb" | "vault"
	...
}

#dockerGateway: #gateway & {
	type:     "docker"
	registry: string
	insecure: bool | *false
	httpOnly: bool | *false
}

#fileGateway: #gateway & {
	type:    "file"
	rootDir: string
}

#ftpGateway: #gateway & {
	type:     "ftp"
	host:     string
	port:     number
	username: string | *""
	password: string | *""
	basePath: string | *""
	insecure: bool | *false
}

#gitGateway: #gateway & {
	type:       "git"
	repoURL:    string
	username:   string | *""
	password:   string | *""
	sshKey:     string | *""
	sshKeyFile: string | *""
	insecure:   bool | *false
	_checks: {
		onlyOneOfPasswordSSHKeyANDSSHKeyFile: true
		onlyOneOfPasswordSSHKeyANDSSHKeyFile: (password == "" && sshKey == "" && sshKeyFile == "") || (password != "" && sshKey == "" && sshKeyFile == "") || (password == "" && sshKey != "" && sshKeyFile == "") || (password == "" && sshKey == "" && sshKeyFile != "")
	}
}

#httpGateway: #gateway & {
	type:     "http"
	baseURL:  string
	username: string | *""
	password: string | *""
	insecure: bool | *false
}

#nfsGateway: #gateway & {
	type: "nfs"
	host: string
	// username: string // TODO: add username and password
	// password: string
}

#s3Gateway: #gateway & {
	type:            "s3"
	endpoint:        string
	accessKeyID:     string
	secretAccessKey: string
	useSSL:          bool | *true
}

#sftpGateway: #gateway & {
	type:       "sftp"
	host:       string
	port:       number
	username:   string
	password:   string | *""
	insecure:   bool
	basePath:   string | *""
	hostKey:    string | *""
	sshKey:     string | *""
	sshKeyFile: string | *""
	insecure:   bool | *false
	_checks: {
		onlyOneOfPasswordSSHKeyANDSSHKeyFile: true
		onlyOneOfPasswordSSHKeyANDSSHKeyFile: (password == "" && sshKey == "" && sshKeyFile == "") || (password != "" && sshKey == "" && sshKeyFile == "") || (password == "" && sshKey != "" && sshKeyFile == "") || (password == "" && sshKey == "" && sshKeyFile != "")
		// when not insecure: hostKey is required
		hostKeyRequired: !insecure
		hostKeyRequired: hostKey != ""
	}
}

#smbGateway: #gateway & {
	type:     "smb"
	host:     string
	port:     number
	username: string
	password: string
}

#vaultGateway: #gateway & {
	type:     "vault"
	endpoint: string
	token:    string
	basePath: string | *""
	insecure: bool | *false
}

#gateways: #fileGateway | #httpGateway | #s3Gateway | #dockerGateway | #gitGateway | #ftpGateway | #nfsGateway | #sftpGateway | #smbGateway | #vaultGateway
