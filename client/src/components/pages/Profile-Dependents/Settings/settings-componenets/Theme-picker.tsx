import { ColorPicker } from "@/components/ui/color-picker";
import { useContext, useState } from "react";
import { ColorContext } from "../settings-logic/color/ColorContext";
import { Button } from "@/components/ui/button";

export const ThemePicker = () => {
  const context = useContext(ColorContext);
  if (!context) {
    throw new Error("ThemePicker must be used within a ColorProvider");
  }

  const {
    colorLeft,
    colorMiddle,
    colorRight,
    setColorLeft,
    setColorMiddle,
    setColorRight,
  } = context;

  // Temporary state for staging changes
  const [tempColorLeft, setTempColorLeft] = useState(colorLeft);
  const [tempColorMiddle, setTempColorMiddle] = useState(colorMiddle);
  const [tempColorRight, setTempColorRight] = useState(colorRight);

  const handleApplyChanges = () => {
    setColorLeft(tempColorLeft);
    setColorMiddle(tempColorMiddle);
    setColorRight(tempColorRight);
  };

  return (
    <div className="w-100 flex items-center justify-center gap-10">
      <ColorPicker
        className="w-10"
        onChange={(v) => {
          setTempColorLeft(v);
        }}
        value={tempColorLeft}
      />
      <ColorPicker
        className="w-10"
        onChange={(v) => {
          setTempColorMiddle(v);
        }}
        value={tempColorMiddle}
      />
      <ColorPicker
        className="w-10"
        onChange={(v) => {
          setTempColorRight(v);
        }}
        value={tempColorRight}
      />
      <Button className="cursor-pointer" onClick={handleApplyChanges}>
        Apply changes
      </Button>
    </div>
  );
};
