import { Button } from "@/components/ui/button";
import { Github, ExternalLink, Mail, Heart } from "lucide-react";

const FooterSection = () => {
  return (
    <footer className="py-16 border-t border-border bg-muted/10">
      <div className="container px-6 mx-auto">
        <div className="grid md:grid-cols-3 gap-8 mb-8">
          {/* Brand */}
          <div className="space-y-4">
            <h3 className="text-2xl font-bold">
              <span className="text-primary">Dev</span>syringe
            </h3>
            <p className="text-muted-foreground leading-relaxed">
              Eliminate manual configuration updates. Inject dynamic values into static files 
              with declarative YAML configs.
            </p>
            <div className="flex space-x-3">
              <Button variant="outline" size="sm" asChild>
                <a href="https://github.com/alchemmist/devsyringe" target="_blank" rel="noopener noreferrer">
                  <Github className="h-4 w-4 mr-2" />
                  GitHub
                </a>
              </Button>
              <Button variant="outline" size="sm" asChild>
                <a href="mailto:anton.ingrish@gmail.com">
                  <Mail className="h-4 w-4 mr-2" />
                  Contact
                </a>
              </Button>
            </div>
          </div>

          {/* Resources */}
          <div className="space-y-4">
            <h4 className="font-semibold text-lg">Resources</h4>
            <div className="space-y-2">
              <a 
                href="https://github.com/alchemmist/devsyringe#readme" 
                target="_blank" 
                rel="noopener noreferrer"
                className="block text-muted-foreground hover:text-foreground transition-colors"
              >
                Documentation
              </a>
              <a 
                href="https://github.com/alchemmist/devsyringe/tree/main/examples" 
                target="_blank" 
                rel="noopener noreferrer"
                className="block text-muted-foreground hover:text-foreground transition-colors"
              >
                Examples
              </a>
              <a 
                href="https://github.com/alchemmist/devsyringe/releases" 
                target="_blank" 
                rel="noopener noreferrer"
                className="block text-muted-foreground hover:text-foreground transition-colors"
              >
                Releases
              </a>
              <a 
                href="https://alchemmist.xyz/articles/the-devsyringe/" 
                target="_blank" 
                rel="noopener noreferrer"
                className="block text-muted-foreground hover:text-foreground transition-colors"
              >
                Article
              </a>
            </div>
          </div>

          {/* Community */}
          <div className="space-y-4">
            <h4 className="font-semibold text-lg">Community</h4>
            <div className="space-y-2">
              <a 
                href="https://github.com/alchemmist/devsyringe/issues" 
                target="_blank" 
                rel="noopener noreferrer"
                className="block text-muted-foreground hover:text-foreground transition-colors"
              >
                Issues & Bug Reports
              </a>
              <a 
                href="https://github.com/alchemmist/devsyringe/discussions" 
                target="_blank" 
                rel="noopener noreferrer"
                className="block text-muted-foreground hover:text-foreground transition-colors"
              >
                Discussions
              </a>
              <a 
                href="https://github.com/alchemmist/devsyringe/blob/main/CONTRIBUTING.md" 
                target="_blank" 
                rel="noopener noreferrer"
                className="block text-muted-foreground hover:text-foreground transition-colors"
              >
                Contributing
              </a>
            </div>
          </div>
        </div>

        {/* Bottom bar */}
        <div className="pt-8 border-t border-border flex flex-col md:flex-row justify-between items-center space-y-4 md:space-y-0">
          <div className="text-sm text-muted-foreground">
            Built by{" "}
            <a 
              href="https://github.com/alchemmist" 
              target="_blank" 
              rel="noopener noreferrer"
              className="text-primary hover:text-primary-glow transition-colors"
            >
              @alchemmist
            </a>
          </div>
          
          <div className="text-sm text-muted-foreground">
            Open source under MIT License
          </div>
        </div>
      </div>
    </footer>
  );
};

export default FooterSection;
