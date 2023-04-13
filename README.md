Real-Time Locator API
This API is for a locator app that allows users to upload their location in real-time and see other users' locations in real-time. It is built using GoLang and runs on an AWS EC2 instance.

Installation
Clone this repository using the command git clone https://github.com/CharlesMuchogo/locator/master
Navigate to the cloned directory using the command cd locator.
Install the dependencies using the command go mod download.
Rename the config-sample.yaml file to config.yaml.
Update the config.yaml file with your desired configurations.
Run the server using the command go run main.go.

Deployment
The API is deployed on an AWS EC2 instance. GitHub Actions is used to create a Docker image of the API and push it to DockerHub. The EC2 instance then pulls the latest Docker image and runs the API server.

Contributing
Feel free to submit pull requests or report issues. For major changes, please open an issue first to discuss what you would like to change.

License
This project is licensed under the MIT License.
