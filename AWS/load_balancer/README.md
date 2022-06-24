<!-- markdownlint-disable -->

# Elastic Load Balancer (ELB)

Created a load Balancer.

<ul>
    <li>Created two new security groups:
        <ol>
            <li>loadbalancer-sg: allowing traffic from **anywhere** type HTTP protocol TCP Port 80</li>
            <li>webtier-sg: allowing traffic from the **anywhere** type SSH protocol TCP Port 22, traffic from the loadbalancer-sg (using its ID) type HTTP protocol TCP port 80, and traffic from webtier-sg type MYSQL/Aurora protocol TCP port 3306</li> 
        </ol></li>
    <li>Created a target group, making sure the AZ is the one where our instance is and <strong>registered</strong> the instance</li>
</ul>
