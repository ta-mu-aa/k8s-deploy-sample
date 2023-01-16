# k8s-deploy-sample

### /go-appをkubernetesで動かすコマンド
1. ディレクトリを移動  
`cd go-app`

1. Dockerイメージを作成  
`docker build ./ -t go-sample-image`

1. デプロイを作成  
`kubectl apply -f deployment.yaml`

1. サービスを作成し、APIにアクセスできるようにする  
 `kubectl apply -f service.yaml`

