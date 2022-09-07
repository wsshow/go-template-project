# ϵͳ����
$arrOS = "linux", "windows", "android"
# �ܹ�����
$arrARCH = "386", "amd64", "arm", "arm64", "riscv", "riscv64"
Write-Host "=========================================="
Write-Host "   [cmd]:   build.ps1 <os> <arch>"
Write-Host "   [os]:   ", $( $arrOS -join "," )
Write-Host "   [arch]: ", $( $arrARCH -join "," )
Write-Host "=========================================="

$os = $args[0]
# ������
if (!$arrOS.Contains($os))
{
    Write-Host "Error: Invalid GOOS" -ForegroundColor:Red;
    exit(-1)
}

$arch = $args[1]
# ������
if (!$arrARCH.Contains("$arch"))
{
    Write-Host $arch, $args[1]
    Write-Host $arrARCH
    Write-Host "Error: Invalid GOARCH" -ForegroundColor:Red;
    exit(-1)
}

# �޸Ļ�������
Write-Output("GOOS $( go env GOOS ) => $os")
$env:GOOS = $os
Write-Output("GOARCH $( go env GOARCH ) => $arch")
$env:GOARCH = $arch

# ��ȡ����汾
$Version = "1.0.1"
# ��ȡ��ǰʱ��
$BuildTime = (Get-Date).ToString("yyyy-MM-dd HH:mm:ss")
# �����ϱ������л��� LDFlags ������
$LDFlags = " \
    -X 'gtp/version.version=$Version' \
    -X 'gtp/version.buildTime=$BuildTime' \
"

# ����
$output = ".\release\$os\gtp-$os-$arch"

if (!(Test-Path -Path ".\release\$os"))
{
    New-Item -ItemType Directory -Path ".\release\$os"
}

if ($os -eq "linux")
{
    go build -ldflags "$LDFlags" -o $output .\main.go
}
if ($os -eq "windows")
{
    $exePath = "$output.exe"
    go build  -ldflags "$LDFlags" -o $exePath .\main.go
}

Write-Host "build had finished, output:$output" -ForegroundColor:Green
# ��explore
Invoke-Item ".\release\$os\"