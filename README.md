# go_serverless
Women Who Go - AWS Serverless 배포

# 시작
## 1. AWS Serverless로 Go 쉽게 배포하기
- 핸즈온 사전 응답: https://docs.google.com/forms/d/e/1FAIpQLSdQmZbMW3A53L3RnSgxCwM9MDiTpk0xs-OeUcczXj9h2Pr8VQ/viewform
- AWS 핸즈온 등록: https://catalog.us-east-1.prod.workshops.aws/sign-in?redirect=%2Fjoin%3Faccess-code%3D711c-032246-c6

### 배포 방법
- 서버 내에서 코드를 직접 배포(state) & 도커로 배포(state less)
- ZIP 파일 압축 후 AWS 람다로 배포(server less)
    - 서버를 생성/관리하지 않고도 애플리케이션을 빌드할 수 있게 해준다.
    - AWS console에서 클릭을 몇번 하여 배포해야 함(cli 도 가능하긴함)
    - zip 파일 압축 및 올리는 것이 번거롭다.
    - 하루 1번의 요청이 올 경우 월비용이 무료이다.
- AWS 서버리스로 배포
    - 서버리스: 클라우드에 API 배포를 쉽게 할 수 있는 프레임워크
    - AWS 외에도 GCP, Azure 등의 클라우드 업체에도 있다.
    - 동시성 설정이 가능하다. (인프라 자동 확장 및 축소)
    - 요청이 없어 호출이 없으면 비용이 발생하지 않는다.
    - 코드와 인프라를 함께 배포하여 바로 사용할 수 있는 서버리스 애플리케이션을 제공하며 여러가지 언어와 상태 등을 지원한다.

## AWS Serverless 를 사용하는 이유
- Serverless 모든 기능이 AWS 에서는 다 가능하다.
- 프로젝트 비용을 줄일 수 있으며, 언제라도 인프라 환경을 삭제 및 재생성 가능하며, 여러 클라우드 환경에서 사용 가능하다.
- 5분 내로 배포를 해볼 수 있다.

## Go
- 정적 컴파일 언어로 미리 컴파일을 해둬서 속도가 빠르며, 타입 에러를 컴파일 시점에 발견할 수 있어 타입 안정성이 뛰어나다. Java도 정적 컴파일 언어인데 Go 보다 느리다고 한다.

# 실행
## 1. CLI export
- AWS account access - Get AWS CLI credentials - 운영체제에 맞는 export cli 복사 - 커맨드 창에 붙여넣음
## 2. Go와 npm 설치
- go: https://go.dev/doc/install
- npm: https://www.npmjs.com/package/npm
## 3. serverless 설치
- npm install -g serverless
## 4. 프로젝트 생성
- serverless create -t aws-go-mod -p go-serverless
## 4-2. 안된다면 이처럼
- serverless create -t aws-go-mod -p myservice-test
- cd myservice-test
- go mod download github.com/aws/aws-lambda-go
- make
## 5-1. aws 람다 다운로드
- go mod download github.com/aws/aws-lambda-go
## 5-2. 고 프로젝트 빌드
- make
## 5-3. 서버리스 배포
- sls deploy

## end. 서버리스 삭제
- sls remove

## 도메인 연결
- 참고: https://docs.aws.amazon.com/ko_kr/apigateway/latest/developerguide/how-to-custom-domains.html
- https://repost.aws/ko/knowledge-center/custom-domain-name-amazon-api-gateway

## 특징
* 배포를 하면 접근 가능한 주소가 나온다. 해당 주소를 curl -XGET 명령어로 호출하면 json 이 응답된다.
* 이 프로세스의 장점: 빠르게 배포 가능하고 인프라 및 엔드포인트를 생성하기 때문에 좋다. 
- 서버리스 참고: https://github.com/serverless/examples
 - Go로 Google Map API 만들기 등의 템플릿이 있다고 한다.











