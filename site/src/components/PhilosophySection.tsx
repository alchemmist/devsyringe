import { Card } from "@/components/ui/card";
import { ArrowRight, Terminal, FileText, Zap } from "lucide-react";

const PhilosophySection = () => {
  return (
    <section className="py-24 relative">
      <div className="container px-6 mx-auto">
        {/* Section header */}
        <div className="text-center mb-16">
          <h2 className="text-3xl md:text-5xl font-bold mb-6">
            The <span className="text-primary">Injection</span> Philosophy
          </h2>
          <p className="text-xl text-muted-foreground max-w-3xl mx-auto">
            Treat static files as dynamic endpoints. One command, universal automation, 
            seamless integration without changing your codebase.
          </p>
        </div>

        {/* Philosophy cards */}
        <div className="grid md:grid-cols-3 gap-8 mb-16">
          <Card className="p-8 bg-card/50 backdrop-blur-sm border-border/50 hover:border-primary/30 transition-all duration-300 group">
            <div className="w-12 h-12 bg-gradient-primary rounded-lg flex items-center justify-center mb-6 group-hover:animate-pulse-glow">
              <Terminal className="h-6 w-6 text-primary-foreground" />
            </div>
            <h3 className="text-xl font-semibold mb-4">Command → Value</h3>
            <p className="text-muted-foreground leading-relaxed">
              Run any CLI command and extract values using regex patterns. No complex parsing, 
              just simple pattern matching that works universally.
            </p>
          </Card>

          <Card className="p-8 bg-card/50 backdrop-blur-sm border-border/50 hover:border-primary/30 transition-all duration-300 group">
            <div className="w-12 h-12 bg-gradient-injection rounded-lg flex items-center justify-center mb-6 group-hover:animate-injection-flow">
              <ArrowRight className="h-6 w-6 text-white" />
            </div>
            <h3 className="text-xl font-semibold mb-4">Value → Files</h3>
            <p className="text-muted-foreground leading-relaxed">
              Inject extracted values into multiple target files simultaneously. 
              Use contextual clues to find the right places for replacement.
            </p>
          </Card>

          <Card className="p-8 bg-card/50 backdrop-blur-sm border-border/50 hover:border-primary/30 transition-all duration-300 group">
            <div className="w-12 h-12 bg-accent/20 rounded-lg flex items-center justify-center mb-6 group-hover:shadow-[0_0_20px_hsl(var(--accent)/0.4)]">
              <Zap className="h-6 w-6 text-accent" />
            </div>
            <h3 className="text-xl font-semibold mb-4">Zero Integration</h3>
            <p className="text-muted-foreground leading-relaxed">
              Your application remains completely unaware. No imports, no hooks, 
              no changes to your code. Pure automation that leaves no trace.
            </p>
          </Card>
        </div>

        {/* Code visualization */}
        <div className="max-w-6xl mx-auto">
          <div className="bg-terminal border border-terminal-border rounded-xl p-8 shadow-terminal">
            <div className="flex items-center mb-6">
              <div className="flex space-x-2">
                <div className="w-3 h-3 bg-destructive rounded-full"></div>
                <div className="w-3 h-3 bg-yellow-500 rounded-full"></div>
                <div className="w-3 h-3 bg-primary rounded-full"></div>
              </div>
              <span className="ml-4 text-sm text-muted-foreground font-mono">devsyringe.yaml</span>
            </div>
            
            <pre className="font-mono text-sm overflow-x-auto">
              <code className="text-foreground whitespace-pre">
{`serums:
  `}<span className="text-accent">localtunnel</span>{`:
    `}<span className="text-primary">source</span>{`: lt --port 8080
    `}<span className="text-primary">mask</span>{`: https://[a-z0-9\\-]+\\.loca\\.lt
    `}<span className="text-primary">targets</span>{`:
      `}<span className="text-accent">env_file</span>{`:
        `}<span className="text-primary">path</span>{`: ./.env
        `}<span className="text-primary">clues</span>{`: [`}<span className="text-yellow-400">"API_BASE_URL"</span>{`]
      `}<span className="text-accent">frontend_config</span>{`:
        `}<span className="text-primary">path</span>{`: ./src/config.js
        `}<span className="text-primary">clues</span>{`: [`}<span className="text-yellow-400">"API_URL"</span>{`, `}<span className="text-yellow-400">"const"</span>{`]`}
              </code>
            </pre>
            
            <div className="mt-6 flex items-center text-muted-foreground text-sm">
              <Terminal className="h-4 w-4 mr-2" />
              <span className="font-mono">dsy inject</span>
              <ArrowRight className="h-4 w-4 mx-2 text-primary" />
              <span>Automatic injection across all targets</span>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
};

export default PhilosophySection;