# cd ..
go build -o ./cmd/shortener/shortener.exe cmd/shortener/main.go
$ITERBRANCH = git rev-parse --abbrev-ref HEAD
if ($ITERBRANCH -match '^iter([0-9]+)$' -eq 'True') {
$ITERNUM = $ITERBRANCH -replace '^iter([0-9]+)$', '$1'
for (($i = 1); $i -le $ITERNUM; $i++)
{
if ($i -le 3){ .\shortenertestbeta.exe --test.v --test.run=^TestIteration$i$ --binary-path=./cmd/shortener/shortener.exe --source-path=./ }
elseif ($i -le 4){ .\shortenertestbeta.exe --test.v --test.run=^TestIteration$i$ --binary-path=./cmd/shortener/shortener.exe --source-path=./ -server-port="8081"}

}
}
Write-Host -NoNewLine 'Press any key to continue...';
$null = $Host.UI.RawUI.ReadKey('NoEcho,IncludeKeyDown');