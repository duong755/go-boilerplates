## 2021/05/14

```shell
dotnet new sln --name OpenLMS

dotnet new api --name OpenLMS.Server --output ./src

dotnet sln add ./src/OpenLMS.Server.csproj

dotnet add ./src/OpenLMS.Server.csproj package Newtonsoft.Json
dotnet add ./src/OpenLMS.Server.csproj package Microsoft.VisualStudio.Web.CodeGeneration.Design
dotnet add ./src/OpenLMS.Server.csproj package Microsoft.EntityFrameworkCore.SqlServer
dotnet add ./src/OpenLMS.Server.csproj package Microsoft.AspNetCore.Identity
dotnet add ./src/OpenLMS.Server.csproj package Swashbuckle.AspNetCore
```

Run server:

```shell
dotnet run --project ./src/OpenLMS.Server.csproj

~/.ngrok/bin/ngrok authtoken <TOKEN> # https://dashboard.ngrok.com/get-started/setup
~/.ngrok/bin/ngrok http -host-header=localhost https://localhost:5001
```

