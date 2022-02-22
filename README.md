## Securely storing environment variables in Google Cloud Functions using yaml

### The action will help you to solve the following problems:
1. If you need to store API keys or other sensitive information in your .yaml as environment variables for deployment on Google Cloud Function.
2. If you don't like the idea to push secret environment variables your .yaml to GitHub.
3. If you don't like the idea to store the environment variables in a datastore.


### Action swaps environment variables in app.yaml with the minimal effort

1. Modify your < NAME >.yaml file:

        KEY1: $KEY1
        KEY2: $KEY2

2. Add this action to your workflow:

        - uses: kiwicampus/gcp-env-vars-file-compiler@main
          with:
            file: ./< NAME >.yaml
          env:
            KEY1: ${{ secrets.KEY1 }}
            KEY2: ${{ secrets.KEY2 }}

Full example with deployment to Google App Engine:     

    deploy:
        name: Deploy
        runs-on: ubuntu-latest
        needs: [build]
    steps:
        - uses: actions/checkout@v2
        - uses: kiwicampus/gcp-env-vars-file-compiler@main
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
