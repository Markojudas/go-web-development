<!-- markdownlint-disable -->

# Exercise : Having our web app highly available while interacting with our database

In this exercise we make our web application, code seen on `main.go` highly available with AWS. The steps are as following:

<ul>
    <li>Create the binary:
        <code>GOOS=linux GOARCH=amd64 go build -o dbExercise</code>
    </li>
    <li>Upload the code to your EC2 Instance:
        <code>scp -i [path-to-pem-file] dbExercise ubuntu@[EC2-instance-public-dns]:</code>
    </li>
    <li>Connect to your EC2 Instance:
        <code>scp -i [path-to-pem-file] ubuntu@[EC2-instance-public-dns]</code>
    </li>
    <li>Change Permissions:
        <code>sudo chmod 777 dbExercise</code>
    </li>
    <li>Create a file on <code>/var/systemd/system</code> directory and give it an arbitrary name but end with <code>.service</code></li>
    <ul>
    <li>code for the file:</li>   
<pre>
<code>[Unit]    
Description=Go Server

<p>[Service]
ExecStart=/home/ubuntu/dbExercise
WorkingDirectory=/home/ubuntu
User=root
Group=root
Restart=always</p>
<p>[Install]
WantedBy=multi-user.target</p></code></pre></ul>
    <li>Enable and start the service</li>
       <p><code>sudo systemctl enable randomname.service</code></p>
       <p><code>sudo systemctl start randomname.service</code></p>
    <li>Create an image of your EC2 Instance (AMI)</li>
    <li>Launch an Instance using the newly created (AMI), making sure it's in another availability zone</li>
    <li>Add both instances to the Target Group</li>
    <li>Connect using the the DNS-NAME of your loadbalancer</li>
</ul>

This assumes the load balancer and security groups are already set up.
