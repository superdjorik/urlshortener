# cd ..
go build -o ./cmd/shortener/shortener.exe cmd/shortener/main.go
$ITERBRANCH = git rev-parse --abbrev-ref HEAD
if ($ITERBRANCH -match '^iter([0-9]+)$' -eq 'True') {
$ITERNUM = $ITERBRANCH -replace '^iter([0-9]+)$', '$1'
for (($i = 0); $i -le $ITERNUM; $i++)
{.\shortenertestbeta.exe --test.v --test.run=^TestIteration$ITERNUM$ --binary-path=./cmd/shortener/shortener.exe --source-path=./}
}
Write-Host -NoNewLine 'Press any key to continue...';
$null = $Host.UI.RawUI.ReadKey('NoEcho,IncludeKeyDown');