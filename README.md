# export-file

Export the output of the web application

## Author

* [@smslash](https://github.com/smslash)

## Run 

```bash
go run ./cmd/web
```

or 

```bash
docker build -t image_name .
```

Here `image_name` is the name you want to assign to your image. The dot at the end of the command points to the current directory where the Dockerfile should be located.

```bash
docker run -d -p 4000:4000 image_name
```

Then open link [http://127.0.0.1:4000](http://127.0.0.1:4000) in your browser.

## Objectives

You must follow the same [principles](https://01.alem.school/git/root/public/src/branch/master/subjects/ascii-art-web/README.md) as the first subject.

Ascii-art-web-export consists on making sure that it is possible to export the output of the web application, at least in one export format at your choice.

- You should be able to export the result of the [ascii-art](https://01.alem.school/git/root/public/src/branch/master/subjects/ascii-art/README.md) project implemented in the website.
- The file must be exported with the right permissions (**read and write**) for the user.

The *Hypertext Transfer Protocol* (HTTP) is the foundation of data communication for the World Wide Web (www). When sending a file or image as part of the HTTP response we must include the use of [HTTP headers](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers).

Headers used for file transfer are [Content-Type](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Type), [Content-Length](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Length) and [Content-Disposition](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Disposition) and more if needed, but for this project you will be obligated to use those headers.

## Instructions

As you already learned from the previous project you must create a new HTTP endpoint to be able to transfer the file to the client. The following instructions must also be followed :

- The web server must export at least in one export format.
- The web server must be created in **Go**.
- The web site must include a button or a link to download/export the file.
- You must handle website errors.
- The code must respect the [good practices](https://01.alem.school/git/root/public/src/branch/master/subjects/good-practices/README.md).

## Allowed packages

- Only the [standard go](https://pkg.go.dev/std) packages are allowed

This project will help you learn about :

- Text File (txt)
- Here are some more [examples](https://en.wikipedia.org/wiki/Document_file_format)
- [HTTP headers](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers)
- Ways to receive data.
- Ways to output data.
