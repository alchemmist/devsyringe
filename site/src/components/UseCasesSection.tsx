import { Card } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Globe, Database, FileText, Cloud, GitBranch, Shield } from "lucide-react";

const UseCasesSection = () => {
  const useCases = [
    {
      icon: Globe,
      title: "Tunnel Automation",
      description: "Automatically inject ngrok, localtunnel, or serveo URLs into your environment files and frontend configs.",
      examples: ["API endpoint updates", "Webhook URL injection", "Development server exposure"],
      color: "text-accent",
      bgColor: "bg-accent/10"
    },
    {
      icon: Database,
      title: "Infrastructure Management", 
      description: "Fetch credentials from HashiCorp Vault, update deployment configs with fresh endpoints and IPs.",
      examples: ["Database credentials", "Service discovery", "Load balancer IPs"],
      color: "text-primary",
      bgColor: "bg-primary/10"
    },
    {
      icon: FileText,
      title: "Dynamic Documentation",
      description: "Keep documentation, résumés, and reports current with live statistics and real-time data.",
      examples: ["GitHub star counts", "API version numbers", "Performance metrics"],
      color: "text-purple-400",
      bgColor: "bg-purple-400/10"
    },
    {
      icon: Cloud,
      title: "CI/CD Integration",
      description: "Inject build metadata, deployment URLs, and environment-specific configurations during pipeline execution.",
      examples: ["Build timestamps", "Docker image tags", "Environment URLs"],
      color: "text-orange-400",
      bgColor: "bg-orange-400/10"
    },
    {
      icon: GitBranch,
      title: "Version Management",
      description: "Automatically update version numbers, release notes, and changelog entries across multiple files.",
      examples: ["Package.json versions", "API documentation", "Release manifests"],
      color: "text-blue-400",
      bgColor: "bg-blue-400/10"
    },
    {
      icon: Shield,
      title: "Security Rotation",
      description: "Rotate API keys, tokens, and secrets across configuration files with automated secret management.",
      examples: ["API key rotation", "Certificate updates", "Token refresh"],
      color: "text-red-400",
      bgColor: "bg-red-400/10"
    }
  ];

  return (
    <section className="py-24 relative">
      <div className="container px-6 mx-auto">
        {/* Section header */}
        <div className="text-center mb-16">
          <h2 className="text-3xl md:text-5xl font-bold mb-6">
            Real-World <span className="text-primary">Use Cases</span>
          </h2>
          <p className="text-xl text-muted-foreground max-w-3xl mx-auto">
            From tunnel URLs to infrastructure automation, Devsyringe eliminates manual 
            configuration updates across your entire development workflow.
          </p>
        </div>

        {/* Use cases grid */}
        <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
          {useCases.map((useCase, index) => {
            const IconComponent = useCase.icon;
            return (
              <Card 
                key={index}
                className="p-6 bg-card/50 backdrop-blur-sm border-border/50 hover:border-primary/30 transition-all duration-300 group hover:animate-float"
              >
                <div className={`w-12 h-12 ${useCase.bgColor} rounded-lg flex items-center justify-center mb-6 group-hover:scale-110 transition-transform`}>
                  <IconComponent className={`h-6 w-6 ${useCase.color}`} />
                </div>
                
                <h3 className="text-xl font-semibold mb-3">{useCase.title}</h3>
                <p className="text-muted-foreground mb-4 leading-relaxed">
                  {useCase.description}
                </p>
                
                <div className="flex flex-wrap gap-2">
                  {useCase.examples.map((example, exampleIndex) => (
                    <Badge 
                      key={exampleIndex}
                      variant="outline" 
                      className="text-xs border-border/50 bg-muted/30"
                    >
                      {example}
                    </Badge>
                  ))}
                </div>
              </Card>
            );
          })}
        </div>

        {/* Featured example */}
        <div className="mt-16 max-w-4xl mx-auto">
          <div className="bg-terminal border border-terminal-border rounded-xl p-8 shadow-terminal">
            <div className="flex items-center mb-6">
              <Globe className="h-5 w-5 text-accent mr-3" />
              <h3 className="text-lg font-semibold">Featured: Tunnel + API Integration</h3>
            </div>
            
            <div className="grid md:grid-cols-2 gap-6">
              <div>
                <div className="text-sm text-muted-foreground mb-3 font-mono">Terminal Output</div>
                <div className="bg-code-bg rounded-lg p-4 font-mono text-sm">
                  <div className="text-accent">$ ssh -R 80:localhost:3000 serveo.net</div>
                  <div className="text-primary mt-1">Forwarding HTTP traffic from</div>
                  <div className="text-primary">https://abc-123.serveo.net</div>
                  <div className="text-muted-foreground">HTTP request from 203.0.113.1</div>
                </div>
              </div>
              
              <div>
                <div className="text-sm text-muted-foreground mb-3 font-mono">Automatic Updates</div>
                <div className="space-y-3">
                  <div className="bg-code-bg rounded-lg p-3 font-mono text-xs">
                    <div className="text-muted-foreground">.env</div>
                    <div className="text-primary">API_URL=https://abc-123.serveo.net</div>
                  </div>
                  <div className="bg-code-bg rounded-lg p-3 font-mono text-xs">
                    <div className="text-muted-foreground">config.js</div>
                    <div className="text-primary">const url = "https://abc-123.serveo.net"</div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
};

export default UseCasesSection;