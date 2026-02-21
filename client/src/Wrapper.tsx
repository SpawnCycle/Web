import { useContext } from "react";
import { ColorContext } from "./components/pages/Profile-Dependents/Settings/settings-logic/color/ColorContext";

interface WrapperProps {
  children: React.ReactNode;
}

export function Wrapper({ children }: WrapperProps) {
  const context = useContext(ColorContext);
  const colorLeft = context?.colorLeft || "#374151";
  const colorMiddle = context?.colorMiddle || "#000000";
  const colorRight = context?.colorRight || "#374151";

  return (
    <div
      className="text-white w-screen h-screen absolute top-0 left-0 flex items-center justify-center"
      style={{
        background: `linear-gradient(to right, ${colorLeft}, ${colorMiddle}, ${colorRight})`,
      }}
    >
      {children}
    </div>
  );
}
