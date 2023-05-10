# SRD: (Go) Static Race Detector

A proof of concept work of extracting Structural Operational Semantics rules and applying it to do static race detection.

## Structural Operational Semantics Rules

$
\textbf{Goroutine Creation}\ \frac{}{(G, M, C) \xrightarrow{\texttt{go f()}} (G \cup \{\texttt{f()}\}, M, C)}
$

$
\textbf{Memory Modification}\ \frac{\texttt{f()}\ \text{is}\ \texttt{data++}}{(G, M, C) \xrightarrow{\texttt{go f()}} (G - \{\texttt{f()}\}, M[\texttt{data} \rightarrow M(\texttt{data}) + 1], C)}
$

$
\textbf{Channel Send}\ \frac{\texttt{f()}\ \text{is}\ \texttt{done <- true}}{(G, M, C) \xrightarrow{\texttt{go f()}} (G - \{\texttt{f()}\}, M, C[\texttt{done} \rightarrow C(\texttt{done}) \cup \{\texttt{true}\}])}
$

$
\textbf{Channel Receive}\ \frac{\texttt{f()}\ \text{is}\ \texttt{<-done}}{(G, M, C) \xrightarrow{\texttt{go f()}} (G - \{\texttt{f()}\}, M, C[\texttt{done} \rightarrow C(\texttt{done}) - \{\texttt{true}\}])}
$

$
\textbf{Print}\ \frac{\texttt{f()}\ \text{is}\ \texttt{fmt.Println(data)}}{(G, M, C) \xrightarrow{\texttt{go f()}} (G \cup \{\texttt{f()}\}, M, C)}
$

## Run

Using Nix:

```shell
nix run nixpkgs#go -- run -race cmd/srd/main.go -- examples/example_*.go
```

## Report

```shell
nix shell nixpkgs#texlive.combined.scheme-full -c pdflatex docs/main.tex
nix shell nixpkgs#texlive.combined.scheme-full -c bibtex main.aux
nix shell nixpkgs#texlive.combined.scheme-full -c pdflatex docs/main.tex
```

## License

This repository content excluding all submodules is licensed under the [MIT License](license.md), third-party code are subject to their original license.
