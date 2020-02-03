CONTRACT=$1
GO_PKG=$2
FOLDER=$3
echo "*****Generating goWrapper for $CONTRACT with go package description $GO_PKG*****"
stopContainer(){
    docker stop $1    
}
generate()
{
    rm -rf $FOLDER && mkdir $FOLDER && \
    docker run -d -it -v $PWD:/tmp eum602/solcgen:latest && \
    container=`docker ps | grep 'solcgen:latest' | awk '{print $1}'`
    sleep 2s
    docker exec -it $container sh -c 'abigen --sol "/tmp/contracts/'$CONTRACT'.sol" --pkg '$GO_PKG' --out /tmp/'$CONTRACT'.go' &&
    docker cp $container:/tmp/$CONTRACT.go ./$FOLDER && \
    ROOT=$PWD
    cd $FOLDER && chown `whoami`:`whoami` ./$CONTRACT.go && cd $ROOT && \
    rm -rf $CONTRACT.go && \
    echo "Build process was successful, terminating docker builder"
    stopContainer $container
}

container=`docker ps | grep 'solcgen:latest' | awk '{print $1}'` && \
if [ -z "$container" ]
then 
    echo "Starting a new container..." && generate
else
    echo "stopping current container"
    docker stop $container && generate
fi