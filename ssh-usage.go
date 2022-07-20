ssh, err := NewSshClient(
	"goodluck",
	"some-host",
	22,
	"/Users/goodluck/.ssh/id_rsa",
	"pem-password")

if err != nil {
	log.Printf("SSH init error %v", err)
} else {
	output, err := ssh.RunCommand("ls")
	fmt.Println(output)
	if err != nil {
		log.Printf("SSH run command error %v", err)
	}
}