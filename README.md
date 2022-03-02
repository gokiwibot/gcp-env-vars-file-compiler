## Safe storage of environment variables in Google Cloud Functions using file.yaml

_Note: This action is a custom add-on made to enrich [deploy-cloud-functions](https://github.com/google-github-actions/deploy-cloud-functions)_

### The action will help you to solve the following problems:
1. If you need to store API keys or other sensitive information in your .yaml as environment variables for deployment on Google Cloud Function.
2. If you don't like the idea to push secret environment variables your .yaml to GitHub.
3. If you don't like the idea to store the environment variables in a datastore.

<p align="center">
  <img src="https://i.imgur.com/iUgzFep.png">
</p>

### Action swaps environment variables in your file.yaml with the minimal effort

1. Modify your < NAME >.yaml file:

        KEY1: $KEY1
        KEY2: $KEY2

2. Add this action to your workflow:

        - uses: gokiwibot/gcp-env-vars-file-compiler@v1.3
          with:
            file: ./< NAME >.yaml
          env:
            KEY1: ${{ secrets.KEY1 }}
            KEY2: ${{ secrets.KEY2 }}

Full example with deployment to Google Cloud Function:     

    deploy:
        name: Deploy
        runs-on: ubuntu-latest
        needs: [build]
    steps:
        - uses: actions/checkout@v2
        - uses: gokiwibot/gcp-env-vars-file-compiler@v1.3
          with:
            file: ./< NAME >.yaml
          env:
            KEY1: ${{ secrets.KEY1 }}
            KEY2: ${{ secrets.KEY2 }}
        - uses: 'google-github-actions/auth@v0'
          with:
            credentials_json: '${{ secrets.GCP_SERVICE_ACCOUNT }}'
        - uses: 'google-github-actions/deploy-cloud-functions@v0'
          with:
            name: 'my-function'
            runtime: 'nodejs16'
            env-vars-file: ./< NAME >.yaml


<!-- GETTING STARTED -->
## Development getting started

This is an example of how you may set up the project locally.

### _Option 1_
**Prerequisites**
* Install docker (https://docs.docker.com/engine/install/)
* Install extensions on vscode:
    * [ms-vscode-remote.remote-containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)
* open the repository with vscode and click on open container option


### _Option 2_
**Prerequisites**
* Install go 1.7 (https://go.dev/doc/install)

### Installation

_After you completed any of the previous options_

1. Install dependencies
   ```sh
   go get -d -v
   ```
2. Run the project
   ```js
   go run main.go
   ```


<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/NewFeature`)
3. Commit your Changes (`git commit -m 'Add some NewFeature'`)
4. Push to the Branch (`git push origin feature/NewFeature`)
5. Open a Pull Request
