<!-- markdownlint-disable -->

# Launching a Go Server on AWS

<h2>Launch an EC2 Instance</h2>

For this example I am running a free-tier EC2 instance, running Ubuntu as the OS.

Once you download the `.pem` file, change the permissions:

```bash
sudo chmod 0400 <path-of-key>.pem
```

<h2>Make the Binary</h2>

```bash
GOOS=linux GOARCH=amd64 go build -o helloWorld
```

<h2>SCP</h2>

Wrote the `hello world` program on my local machine and scp to the server:

```bash
scp -i <path-to-pem-file> helloworld ubuntu@<public-DNS-address>:
```

<h2>Connect to Server</h2>

```bash
ssh -i <path-to-pem-file> ubuntu@<public-DNS-addess>
```

<h2>Change Permissions</h2>

The file, named `helloworld` was copied to the home directory.

```bash
sudo chmod 0700 helloworld
```

<h2>Add Persistance</h2>

```bash
cd /etc/systemd/system/
sudo vim helloworld.service
```

Write the following into the file:

```
[Unit]
Description=GoServer

[Service]
ExecStart=/home/ubuntu/helloworld
WorkingDirectory=/home/ubuntu
User=root
Group=root
Restart=always

[Install]
WantedBy=multi-user.target
```

Then run the following commands:

```
sudo systemctl enable helloworld.service
sudo systemctl start hellowrld.service
```

Now you can go on your local machine: `<public-IP>` on your browser!
