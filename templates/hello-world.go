package templates

const (
	webHead = `<html>
  <head>
    <meta http-equiv="refresh" content="5"> 
    <title>Rancher</title>
    <link rel="icon" href="img/favicon.png">
    <style>
      body {
        background-color: white;
        text-align: center;
        padding: 50px;
        font-family: "Open Sans","Helvetica Neue",Helvetica,Arial,sans-serif;
      }
      button {
          background-color: #0075a8; 
          border: none;
          color: white;
          padding: 15px 32px;
          text-align: center;
          text-decoration: none;
          display: inline-block;
          font-size: 16px;
      }

      #logo {
        margin-bottom: 40px;
      }
    </style>
    <script>
      function getCookie(NameOfCookie){
          if (document.cookie.length > 0) {              
          begin = document.cookie.indexOf(NameOfCookie+"=");       
          if (begin != -1) {           
            begin += NameOfCookie.length+1;       
            end = document.cookie.indexOf(";", begin);
            if (end == -1) end = document.cookie.length;
              return unescape(document.cookie.substring(begin, end));
          } 
        }
        return null;
      }

      function setCookie(NameOfCookie, value, expiredays) {
      var ExpireDate = new Date ();
      ExpireDate.setTime(ExpireDate.getTime() + (expiredays * 24 * 3600 * 1000));

        document.cookie = NameOfCookie + "=" + escape(value) + 
        ((expiredays == null) ? "" : "; expires=" + ExpireDate.toGMTString());
      }

      function delCookie (NameOfCookie) {
        if (getCookie(NameOfCookie)) {
          document.cookie = NameOfCookie + "=" +
          "; expires=Thu, 01-Jan-70 00:00:01 GMT";
        }
      }
      function showHeaders() {
          var b = document.getElementById("rancherButton");
          var x = document.getElementById("rancherHeaders");
          if (x.style.display === "none") {
              x.style.display = "block";
              b.innerHTML = "Hide request headers";
              setCookie('viewHeaders','true',365)
          } else {
              x.style.display = "none";
              b.innerHTML = "Show request headers";
              setCookie('viewHeaders','false',365)
          }
      }
      function DoTheCookieStuff()
      {
        viewer=getCookie('viewHeaders');
        if (viewer=='true') {showHeaders()}
      }
    </script>
  </head>
  <body onLoad="DoTheCookieStuff()">
    <img id="logo" src="img/rancher-logo.svg" alt="Rancher logo" width=400 />
    <h1>Hello world!</h1>`

	webDeploy = `    <div id='rancherDeployment'>
      <h3>{{.Deployname}} version {{.Version}}</h3>
      <b>Pod name</b> {{.Podname}}<br />
      <b>Node name</b> {{.Nodename}}<br />
      <b>Host name</b> {{.Host}}<br />
    </div>`

	webServices = `{{- $length := len .Services }} 
  {{- if gt $length 0 }}
    <div id='rancherServices'>
      <h3>Services</h3>
    {{ range $k,$v := .Services }}
      <b>{{ $k }}</b> {{ $v }}<br />
    {{ end }}
    </div>
    <br />
  {{ end }}`

	webHeaders = `    <button id='rancherButton' class='button' onclick='showHeaders()'>Show request headers</button>
    <div id="rancherHeaders" style='display:none'>
      <br />
    {{ range $k,$v := .Headers }}
      <b>{{ $k }}:</b> {{ $v }}<br />
    {{ end }}
    </div>
    <br />`

	webLinks = `    <div id='rancherLinks' class="row social">
      <a class="p-a-xs" href="https://rancher.com/docs"><img src="img/favicon.png" alt="Docs" height="25" width="25"></a>
      <a class="p-a-xs" href="https://slack.rancher.io/"><img src="img/icon-slack.svg" alt="slack" height="25" width="25"></a>
      <a class="p-a-xs" href="https://github.com/rancher/rancher"><img src="img/icon-github.svg" alt="github" height="25" width="25"></a>
      <a class="p-a-xs" href="https://twitter.com/Rancher_Labs"><img src="img/icon-twitter.svg" alt="twitter" height="25" width="25"></a>
      <a class="p-a-xs" href="https://www.facebook.com/rancherlabs/"><img src="img/icon-facebook.svg" alt="facebook" height="25" width="25"></a>
      <a class="p-a-xs" href="https://www.linkedin.com/groups/6977008/profile"><img src="img/icon-linkedin.svg" height="25" alt="linkedin" width="25"></a>
    </div>
    <br />`

	webTail = `  </body>
</html>`

	HelloWorldTemplate = webHead + `
` + webDeploy + `
` + webServices + `
` + webLinks + `
` + webHeaders + `
` + webTail

	HelloWorldDeploy   = webDeploy
	HelloWorldServices = webServices
	HelloWorldHeaders  = webHeaders
)
