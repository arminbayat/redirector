
env > .env


echo "processing build..."
oc create -f buildconfig.yml

echo "creating imageStream..."
oc create imagestream redirector

echo "processing service..."
oc process -f Services.yml --param-file=.env --param SERVICE_ROLE=producer --ignore-unknown-parameters=true | oc apply -f -

echo "processing routes..."
oc process -f Routes.yml --param-file=.env --ignore-unknown-parameters=true | oc apply -f -

echo "processing deployment config..."
oc process -f DeploymentConfig.yml --param-file=.env --ignore-unknown-parameters=true | oc apply -f -;

echo "rolling out $SERVICE_NAME..."
oc rollout latest dc/$SERVICE_NAME
