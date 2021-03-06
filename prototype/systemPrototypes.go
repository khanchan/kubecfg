package prototype

var defaultPrototypes = []*SpecificationSchema{
	&SpecificationSchema{
		APIVersion: "0.1",
		Name:       "io.ksonnet.pkg.namespace",
		Params: ParamSchemas{
			RequiredParam("name", "name", "Name to give the namespace.", String),
		},
		Template: SnippetSchema{
			Description: `A simple namespace. Labels are automatically populated from the name of the
namespace.`,
			ShortDescription: `Namespace with labels automatically populated from the name`,
			YAMLBody: []string{
				"kind: Namespace",
				"apiVersion: v1",
				"metadata:",
				"  name: ${name}",
				"  labels:",
				"    name: ${name}",
			},
			JSONBody: []string{
				`{`,
				`  "kind": "Namespace",`,
				`  "apiVersion": "v1",`,
				`  "metadata": {`,
				`    "name": ${name},`,
				`    "labels": {`,
				`      "name": ${name}`,
				`    }`,
				`  }`,
				`}`,
			},
			JsonnetBody: []string{
				`local k = import "k.libsonnet";`,
				`local ns = k.core.v1.namespace;`,
				``,
				`ns.new() +`,
				`ns.mixin.metadata.name(${name}) +`,
				`ns.mixin.metadata.labels({name: ${name}})`,
			},
		},
	},
	&SpecificationSchema{
		APIVersion: "0.1",
		Name:       "io.ksonnet.pkg.single-port-service",
		Params: ParamSchemas{
			RequiredParam("name", "serviceName", "Name of the service", String),
			RequiredParam("targetLabelSelector", "selector", "Label for the service to target (e.g., 'app: MyApp').", Object),
			RequiredParam("servicePort", "port", "Port for the service to expose.", NumberOrString),
			RequiredParam("targetPort", "port", "Port for the service target.", NumberOrString),
			OptionalParam("protocol", "protocol", "Protocol to use (either TCP or UDP).", "TCP", String),
		},
		Template: SnippetSchema{
			Description: `A service that exposes 'servicePort', and directs traffic
to 'targetLabelSelector', at 'targetPort'.`,
			ShortDescription: `Service that exposes a single port`,
			YAMLBody: []string{
				"kind: Service",
				"apiVersion: v1",
				"metadata:",
				"  name: ${name}",
				"spec:",
				"  selector:",
				"    ${targetLabelSelector}",
				"  ports:",
				"  - protocol: ${protocol}",
				"    port: ${servicePort}",
				"    targetPort: ${targetPort}",
			},
			JSONBody: []string{
				`{`,
				`  "kind": "Service",`,
				`  "apiVersion": "v1",`,
				`  "metadata": {`,
				`    "name": ${name}`,
				`  },`,
				`  "spec": {`,
				`    "selector": {`,
				`      ${targetLabelSelector}`,
				`    },`,
				`    "ports": [`,
				`      {`,
				`        "protocol": ${protocol},`,
				`        "port": ${servicePort},`,
				`        "targetPort": ${targetPort}`,
				`      }`,
				`    ]`,
				`  }`,
				`}`,
			},
			JsonnetBody: []string{
				`local k = import "k.libsonnet";`,
				`local service = k.core.v1.service;`,
				`local port = k.core.v1.service.mixin.spec.portsType;`,
				``,
				`service.new(`,
				`  ${name},`,
				`  ${targetLabelSelector},`,
				`  port.new(${servicePort}, ${targetPort}))`,
			},
		},
	},
	&SpecificationSchema{
		APIVersion: "0.1",
		Name:       "io.ksonnet.pkg.configMap",
		Params: ParamSchemas{
			RequiredParam("name", "name", "Name to give the configMap.", String),
			OptionalParam("data", "data", "Data for the configMap.", "{}", Object),
		},
		Template: SnippetSchema{
			Description:      `A simple config map with optional user-specified data.`,
			ShortDescription: `A simple config map with optional user-specified data`,
			YAMLBody: []string{
				"apiVersion: v1",
				"kind: ConfigMap",
				"metadata:",
				"  name: ${name}",
				"data: ${data}",
			},
			JSONBody: []string{
				`{`,
				`  "apiVersion": "v1",`,
				`  "kind": "ConfigMap",`,
				`  "metadata": {`,
				`    "name": ${name}`,
				`  },`,
				`  "data": ${data}`,
				`}`,
			},
			JsonnetBody: []string{
				`local k = import "k.libsonnet";`,
				`local configMap = k.core.v1.configMap;`,
				``,
				`configMap.new() +`,
				`configMap.mixin.metadata.name("${name}") +`,
				`configMap.data("${data}")`,
			},
		},
	},
	&SpecificationSchema{
		APIVersion: "0.1",
		Name:       "io.ksonnet.pkg.single-port-deployment",
		Params: ParamSchemas{
			RequiredParam("name", "deploymentName", "Name of the deployment", String),
			RequiredParam("image", "containerImage", "Container image to deploy", String),
			OptionalParam("replicas", "replicas", "Number of replicas", "1", Number),
			OptionalParam("port", "containerPort", "Port to expose", "80", NumberOrString),
		},
		Template: SnippetSchema{
			Description: `A deployment that replicates container 'image' some number of times
(default: 1), and exposes a port (default: 80). Labels are automatically
populated from 'name'.`,
			ShortDescription: `Replicates a container n times, exposes a single port`,
			YAMLBody: []string{
				"apiVersion: apps/v1beta1",
				"kind: Deployment",
				"metadata:",
				"  name: ${name}",
				"spec:",
				"  replicas: ${replicas:1}",
				"  template:",
				"    metadata:",
				"      labels:",
				"        app: ${name}",
				"    spec:",
				"      containers:",
				"      - name: ${name}",
				"        image: ${image}",
				"        ports:",
				"        - containerPort: ${port:80}",
			},
			JSONBody: []string{
				`{`,
				`  "apiVersion": "apps/v1beta1",`,
				`  "kind": "Deployment",`,
				`  "metadata": {`,
				`    "name": ${name}`,
				`  },`,
				`  "spec": {`,
				`    "replicas": ${replicas:1},`,
				`    "template": {`,
				`      "metadata": {`,
				`        "labels": {`,
				`          "app": ${name}`,
				`        }`,
				`      },`,
				`      "spec": {`,
				`        "containers": [`,
				`          {`,
				`            "name": ${name},`,
				`            "image": ${image},`,
				`            "ports": [`,
				`              {`,
				`                "containerPort": ${port:80}`,
				`              }`,
				`            ]`,
				`          }`,
				`        ]`,
				`      }`,
				`    }`,
				`  }`,
				`}`,
			},
			JsonnetBody: []string{
				`local k = import "k.libsonnet";`,
				`local deployment = k.apps.v1beta1.deployment;`,
				`local container = k.apps.v1beta1.deployment.mixin.spec.template.spec.containersType;`,
				`local port = container.portsType;`,
				``,
				`deployment.new(`,
				`  ${name},`,
				`  ${replicas},`,
				`  container.new(${name}, ${image}) +`,
				`    container.ports(port.new(${port:80})),`,
				`  {app: ${name}})`,
			},
		},
	},
}
