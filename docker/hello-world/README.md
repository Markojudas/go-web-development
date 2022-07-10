<!-- markdownlint-disable -->

# Build Image from Dockerfile & Run container

<h2>Build Image</h2>

<div class="language-console highlighter-rouge"><div class="highlight"><pre class="highlight"><code><span class="gp">$</span><span class="w"> </span> docker build -t hello-world .
</code></pre></div>    </div>

<h2>Test</h2>
<div class="language-console highlighter-rouge"><div class="highlight"><pre class="highlight"><code><span class="gp">$</span><span class="w"> </span> docker images
</code></pre></div>    </div>
<br>
<br>
<h2>Run Web App on a container</h2>
<div class="language-console highlighter-rouge"><div class="highlight"><pre class="highlight"><code><span class="gp">$</span><span class="w"> </span> docker run -d -p 80:80 hello-world
</code></pre></div>    </div>

<h2>Test</h2>
<div class="language-console highlighter-rouge"><div class="highlight"><pre class="highlight"><code><span class="gp">$</span><span class="w"> </span> docker ps
</code></pre></div>    </div>

go to http://localhost
<br>
<br>

<h2>Stop Container</h2>
<div class="language-console highlighter-rouge"><div class="highlight"><pre class="highlight"><code><span class="gp">$</span><span class="w"> </span> docker stop [container-id]
</code></pre></div>    </div>
