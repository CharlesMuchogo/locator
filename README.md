<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
</head>
<body>
	<h1>Real-Time Locator API</h1>
	<p>This API is for a locator app that allows users to upload their location in real-time and see other users' locations in real-time. It is built using GoLang and runs on an AWS EC2 instance.</p>

<h2>Installation</h2>
<ol>
	<li>Clone this repository using the command <code>git clone https://github.com/CharlesMuchogo/locator/master.</code></li>
	<li>Navigate to the cloned directory using the command <code>cd locator</code>.</li>
	<li>Install the dependencies using the command <code>go mod download</code>.</li>
	<li>Rename the <code>config-sample.yaml</code> file to <code>config.yaml</code>.</li>
	<li>Update the <code>config.yaml</code> file with your desired configurations.</li>
	<li>Run the server using the command <code>go run main.go</code>.</li>
</ol>



<h2>Deployment</h2>
<p>The API is deployed on an AWS EC2 instance. GitHub Actions is used to create a Docker image of the API and push it to DockerHub. The EC2 instance then pulls the latest Docker image and runs the API server.</p>

<h2>Contributing</h2>
<p>Feel free to submit pull requests or report issues. For major changes, please open an issue first to discuss what you would like to change.</p>

</body>
</html>
