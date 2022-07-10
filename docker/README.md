<!-- markdownlint-disable -->

# DOCKER

<h2>Install Docker on Fedora</h2>

<div class="language-console highlighter-rouge"><div class="highlight"><pre class="highlight"><code><span class="gp">$</span><span class="w"> </span><span class="nb">sudo </span>dnf <span class="nt">-y</span> <span class="nb">install </span>dnf-plugins-core
<span class="go">
</span><span class="gp">$</span><span class="w"> </span><span class="nb">sudo </span>dnf config-manager <span class="se">\</span>
    <span class="nt">--add-repo</span> <span class="se">\</span>
    https://download.docker.com/linux/fedora/docker-ce.repo
</code></pre></div></div>

<h2>Install Docker Engine</h2>

<div class="language-console highlighter-rouge"><div class="highlight"><pre class="highlight"><code><span class="gp">$</span><span class="w"> </span><span class="nb">sudo </span>dnf <span class="nb">install </span>docker-ce docker-ce-cli containerd.io docker-compose-plugin
</code></pre></div>    </div>

<h2>Change Permissions</h2>

<div class="language-console highlighter-rouge"><div class="highlight"><pre class="highlight"><code><span class="gp">$</span><span class="w"> </span><span class="nb">sudo </span>usermod -a -G docker $(whoami)
</code></pre></div>    </div>

**you might have to restart your system for the changes to take place**

<h2>Start Docker</h2>

<div class="language-console highlighter-rouge"><div class="highlight"><pre class="highlight"><code><span class="gp">$</span><span class="w"> </span><span class="nb">sudo </span>systemctl start docker
</code></pre></div>    </div>
