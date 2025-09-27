import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { ExternalLink, Github, Download, ArrowRight } from "lucide-react";

const HeroSection = () => {
  return (
    <section className="relative min-h-screen flex items-center justify-center overflow-hidden">
      {/* Animated background grid */}
      <div className="absolute inset-0 bg-[linear-gradient(to_right,hsl(var(--border))_1px,transparent_1px),linear-gradient(to_bottom,hsl(var(--border))_1px,transparent_1px)] bg-[size:4rem_4rem] opacity-20" />
      
      {/* Subtle glow effect */}
      <div className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[600px] h-[300px] bg-gradient-glow opacity-20 rounded-full blur-3xl" />
      
      <div className="container relative z-10 px-6 mx-auto text-center">
        {/* Status badge */}
        <div className="flex justify-center mb-8">
          <Badge variant="outline" className="px-4 py-2 text-sm font-mono border-primary/20 bg-primary/5 text-primary">
            <span className="w-2 h-2 bg-primary/60 rounded-full mr-2" />
            Go CLI Tool • Open Source
          </Badge>
        </div>

        {/* Main heading */}
        <h1 className="text-5xl md:text-7xl font-bold mb-6 leading-tight">
          <span className="text-foreground">Stop</span>{" "}
          <span className="text-muted-foreground">Copy-Pasting</span>
          <br />
          <span className="bg-gradient-injection bg-clip-text text-transparent">
            Dynamic Values
          </span>
        </h1>

        {/* Subtitle */}
        <p className="text-xl md:text-2xl text-muted-foreground mb-8 max-w-3xl mx-auto leading-relaxed">
          Devsyringe automates the tedious workflow of updating static files with dynamic values. 
          <span className="text-accent font-medium"> Inject once, never copy-paste again.</span>
        </p>

        {/* Quick demo visualization */}
        <div className="mb-12 max-w-4xl mx-auto">
          <div className="grid md:grid-cols-3 gap-4 items-center">
            {/* Source */}
            <div className="bg-terminal border border-terminal-border rounded-lg p-4 shadow-terminal">
              <div className="text-sm text-muted-foreground mb-2 font-mono">Source Command</div>
              <div className="font-mono text-sm text-accent">
                $ lt --port 8080
              </div>
              <div className="font-mono text-xs text-primary/80 mt-1">
                https://abc-123.loca.lt
              </div>
            </div>

            {/* Arrow with injection flow */}
            <div className="flex items-center justify-center">
              <div className="w-12 h-1 bg-gradient-injection rounded-full" />
              <ArrowRight className="ml-2 h-6 w-6 text-primary/80" />
            </div>

            {/* Target */}
            <div className="bg-code-bg border border-border rounded-lg p-4">
              <div className="text-sm text-muted-foreground mb-2 font-mono">Target Files</div>
              <div className="font-mono text-xs space-y-1">
                <div className="text-muted-foreground">.env</div>
                <div className="text-primary">API_URL=https://abc-123.loca.lt</div>
                <div className="text-muted-foreground mt-2">config.js</div>
                <div className="text-primary">const url = "https://abc-123.loca.lt"</div>
              </div>
            </div>
          </div>
        </div>

        {/* CTA buttons */}
        <div className="flex flex-col sm:flex-row gap-4 justify-center items-center">
          <Button size="lg" className="group px-8 py-6 text-lg font-medium">
            <Download className="mr-2 h-5 w-5 group-hover:animate-bounce" />
            Install Devsyringe
          </Button>
          
          <Button 
            variant="outline" 
            size="lg" 
            className="px-8 py-6 text-lg group border-muted hover:border-primary"
            asChild
          >
            <a href="https://github.com/alchemmist/devsyringe" target="_blank" rel="noopener noreferrer">
              <Github className="mr-2 h-5 w-5 group-hover:rotate-12 transition-transform" />
              View Source
              <ExternalLink className="ml-2 h-4 w-4 opacity-60" />
            </a>
          </Button>
        </div>

        {/* Quick install */}
        <div className="mt-8 max-w-xl mx-auto">
          <div className="bg-terminal border border-terminal-border rounded-lg p-4 text-left">
            <div className="text-sm text-muted-foreground mb-2 font-mono">Quick Install</div>
            <code className="text-accent font-mono text-sm">
              go install github.com/alchemmist/devsyringe/cmd/dsy@latest
            </code>
          </div>
        </div>
      </div>
    </section>
  );
};

export default HeroSection;