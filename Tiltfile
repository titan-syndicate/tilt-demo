# echo-parent-----------------------------------------------------
k8s_yaml('echo-parent/go-server-deployment.yaml')

docker_build('rianfowler/echo-parent', 'echo-parent')

k8s_resource('echo-parent', port_forwards='3333')

# echo-child-1-----------------------------------------------------
k8s_yaml('echo-child-1/go-server-deployment.yaml')

# (image name, folder)
docker_build('rianfowler/echo-child-1', 'echo-child-1')

# (tilt resource name, port_forwards='localhost:3334')
# k8s_resource('echo-child-1', port_forwards='3334')
k8s_resource('echo-child-1')

# echo-child-2-----------------------------------------------------
k8s_yaml('echo-child-2/go-server-deployment.yaml')

# (image name, folder)
docker_build('rianfowler/echo-child-2', 'echo-child-2')

# (tilt resource name, port_forwards='localhost:3334')
# k8s_resource('echo-child-1', port_forwards='3334')
k8s_resource('echo-child-2')

# echo-grandchild-1-----------------------------------------------------
k8s_yaml('echo-grandchild-1/go-server-deployment.yaml')

# (image name, folder)
docker_build('rianfowler/echo-grandchild-1', 'echo-grandchild-1')

# (tilt resource name, port_forwards='localhost:3334')
# k8s_resource('echo-child-1', port_forwards='3334')
k8s_resource('echo-grandchild-1')