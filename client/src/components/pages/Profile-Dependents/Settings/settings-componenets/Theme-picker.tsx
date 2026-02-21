import { ColorPicker } from "@/components/ui/color-picker";
import { useContext, useState } from "react";
import { ColorContext } from "../settings-logic/color/ColorContext";
import { Button } from "@/components/ui/button";
import { Label } from "@/components/ui/label";
import { useSettings } from "../settings-logic/SettingsContext";

export const ThemePicker = () => {
  const context = useContext(ColorContext);
  if (!context) {
    throw new Error("ThemePicker must be used within a ColorProvider");
  }
  const { settings } = useSettings();

  const {
    colorLeft,
    colorMiddle,
    colorRight,
    setColorLeft,
    setColorMiddle,
    setColorRight,
  } = context;

  // Temporary state for staging changes - synced with applied theme
  const [tempColorLeft, setTempColorLeft] = useState(colorLeft);
  const [tempColorMiddle, setTempColorMiddle] = useState(colorMiddle);
  const [tempColorRight, setTempColorRight] = useState(colorRight);

  // Sync temp colors when applied theme changes
  if (tempColorLeft !== colorLeft) setTempColorLeft(colorLeft);
  if (tempColorMiddle !== colorMiddle) setTempColorMiddle(colorMiddle);
  if (tempColorRight !== colorRight) setTempColorRight(colorRight);

  const handleApplyChanges = () => {
    setColorLeft(tempColorLeft);
    setColorMiddle(tempColorMiddle);
    setColorRight(tempColorRight);
  };

  return (
    <div className="w-100 flex items-center justify-center gap-1">
      <Label
        className={`text-white p-1.5 ${settings.useLiquidGlass ? "[text-shadow:0_2px_4px_rgba(163,163,163,0.8)]" : ""}`}
      >
        Custom Theme
      </Label>
      <ColorPicker
        className="w-10 cursor-pointer"
        onChange={(v) => {
          setTempColorLeft(v);
        }}
        value={tempColorLeft}
      />
      <ColorPicker
        className="w-10 cursor-pointer"
        onChange={(v) => {
          setTempColorMiddle(v);
        }}
        value={tempColorMiddle}
      />
      <ColorPicker
        className="w-10 cursor-pointer"
        onChange={(v) => {
          setTempColorRight(v);
        }}
        value={tempColorRight}
      />
      <Button
        className={`text-white cursor-pointer ${settings.useLiquidGlass ? "bg-white/30 backdrop-blur-lg border-white/30 shadow-sm shadow-white/20[text-shadow:0_2px_4px_rgba(163,163,163,0.8)]" : ""}`}
        onClick={handleApplyChanges}
      >
        Apply changes
      </Button>
    </div>
  );
};
