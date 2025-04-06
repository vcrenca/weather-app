import { Card } from "@weather-api/components/Card";

export default function Home() {
  return (
    <div className="min-h-screen bg-gray-100 p-8 flex items-center justify-center">
      <Card title="Weather App" description="Welcome to the weather app" />
    </div>
  );
}
