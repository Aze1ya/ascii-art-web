run-ascii:
	    docker build --tag ascii-art-web-dockerize .
       
	    docker run -p 8080:8080 -it ascii-art-web-dockerize
       
