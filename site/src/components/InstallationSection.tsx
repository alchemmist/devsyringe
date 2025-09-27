import { Card } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Copy, ExternalLink, Download, Package, Terminal } from "lucide-react";
import { toast } from "sonner";

const InstallationSection = () => {
  const installMethods = [
    {
      title: "Go Install",
      description: "Direct installation via Go toolchain",
      command: "go install github.com/alchemmist/devsyringe/cmd/dsy@latest",
      icon: Download,
      badge: "Recommended",
      badgeVariant: "default" as const
    },
    {
      title: "Homebrew",
      description: "macOS package manager",
      command: "brew tap alchemmist/homebrew-tap && brew install devsyringe", 
      icon: Package,
      badge: "macOS",
      badgeVariant: "secondary" as const
    },
    {
      title: "Arch Linux",
      description: "AUR package manager",
      command: "paru -S devsyringe",
      icon: Package,
      badge: "AUR",
      badgeVariant: "outline" as const
    }
  ];

  const copyToClipboard = (text: string, method: string) => {
    navigator.clipboard.writeText(text);
    toast.success(`${method} command copied to clipboard!`);
  };

  return (
    <section className="py-24 relative bg-muted/20">
      <div className="container px-6 mx-auto">
        {/* Section header */}
        <div className="text-center mb-16">
          <h2 className="text-3xl md:text-5xl font-bold mb-6">
            Get <span className="text-primary">Started</span> in Seconds
          </h2>
          <p className="text-xl text-muted-foreground max-w-2xl mx-auto">
            Multiple installation methods for all platforms. Choose the one that fits your workflow.
          </p>
        </div>

        {/* Installation methods */}
        <div className="grid md:grid-cols-3 gap-6 mb-12">
          {installMethods.map((method, index) => {
            const IconComponent = method.icon;
            return (
              <Card key={index} className="p-6 bg-card border-border hover:border-primary/30 transition-all duration-300 group">
                <div className="flex items-center justify-between mb-4">
                  <div className="flex items-center">
                    <div className="w-10 h-10 bg-primary/10 rounded-lg flex items-center justify-center mr-3">
                      <IconComponent className="h-5 w-5 text-primary" />
                    </div>
                    <h3 className="font-semibold">{method.title}</h3>
                  </div>
                  <Badge variant={method.badgeVariant} className="text-xs">
                    {method.badge}
                  </Badge>
                </div>
                
                <p className="text-sm text-muted-foreground mb-4">{method.description}</p>
                
                <div className="bg-terminal border border-terminal-border rounded-lg p-3 mb-4">
                  <code className="font-mono text-sm text-accent break-all">
                    {method.command}
                  </code>
                </div>
                
                <Button 
                  variant="outline" 
                  size="sm" 
                  className="w-full group-hover:border-primary"
                  onClick={() => copyToClipboard(method.command, method.title)}
                >
                  <Copy className="h-4 w-4 mr-2" />
                  Copy Command
                </Button>
              </Card>
            );
          })}
        </div>

        {/* Quick start guide */}
        <div className="max-w-4xl mx-auto">
          <Card className="p-8 bg-card/50 backdrop-blur-sm border-border/50">
            <h3 className="text-2xl font-semibold mb-6 flex items-center">
              <Terminal className="h-6 w-6 text-primary mr-3" />
              Quick Start Guide
            </h3>
            
            <div className="grid md:grid-cols-2 gap-8">
              <div className="space-y-4">
                <div>
                  <div className="text-sm text-muted-foreground mb-2 font-mono">1. Create config file</div>
                  <div className="bg-terminal border border-terminal-border rounded-lg p-4">
                    <div className="text-sm text-muted-foreground mb-2">devsyringe.yaml</div>
                    <pre className="font-mono text-xs text-foreground">
{`serums:
  tunnel:
    source: lt --port 8080
    mask: https://[a-z0-9\\-]+\\.loca\\.lt
    targets:
      env:
        path: .env
        clues: ["API_URL"]`}
                    </pre>
                  </div>
                </div>
                
                <div>
                  <div className="text-sm text-muted-foreground mb-2 font-mono">2. Run injection</div>
                  <div className="bg-terminal border border-terminal-border rounded-lg p-3">
                    <code className="font-mono text-sm text-accent">dsy inject</code>
                  </div>
                </div>
              </div>
              
              <div className="space-y-4">
                <div>
                  <div className="text-sm text-muted-foreground mb-2 font-mono">3. Monitor processes</div>
                  <div className="bg-terminal border border-terminal-border rounded-lg p-3">
                    <code className="font-mono text-sm text-accent">dsy</code>
                    <div className="text-xs text-muted-foreground mt-1">Opens TUI interface</div>
                  </div>
                </div>
                
                <div>
                  <div className="text-sm text-muted-foreground mb-2 font-mono">4. View logs</div>
                  <div className="bg-terminal border border-terminal-border rounded-lg p-3">
                    <code className="font-mono text-sm text-accent">dsy logs tunnel</code>
                  </div>
                </div>
              </div>
            </div>
            
            <div className="mt-8 flex flex-col sm:flex-row gap-4">
              <Button asChild className="flex-1">
                <a href="https://github.com/alchemmist/devsyringe#readme" target="_blank" rel="noopener noreferrer">
                  <ExternalLink className="h-4 w-4 mr-2" />
                  Full Documentation
                </a>
              </Button>
              <Button variant="outline" asChild className="flex-1">
                <a href="https://github.com/alchemmist/devsyringe/tree/main/examples" target="_blank" rel="noopener noreferrer">
                  <ExternalLink className="h-4 w-4 mr-2" />
                  View Examples
                </a>
              </Button>
            </div>
          </Card>
        </div>
      </div>
    </section>
  );
};

export default InstallationSection;