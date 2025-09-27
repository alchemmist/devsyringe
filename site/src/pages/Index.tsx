import HeroSection from "@/components/HeroSection";
import PhilosophySection from "@/components/PhilosophySection";
import UseCasesSection from "@/components/UseCasesSection";
import InstallationSection from "@/components/InstallationSection";
import FooterSection from "@/components/FooterSection";

const Index = () => {
  return (
    <div className="min-h-screen bg-background">
      <HeroSection />
      <PhilosophySection />
      <UseCasesSection />
      <InstallationSection />
      <FooterSection />
    </div>
  );
};

export default Index;
