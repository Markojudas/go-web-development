<!-- markdownlint-disable -->

# AWS AMI

Creating an AMI (Amazon Machine Image) of our EC2 (Elastic Compute Cloud) Instance to use with the Target Group and the Elastic Load Balancer (ELB).
Launched another EC2 instance using the newly created AMI and then **registered** the instance to the **target group**. This way the ELB directs traffic to either 2 instances!.
