import AccountMenu from "./AccountMenu";
import { Label } from "@radix-ui/react-dropdown-menu";
import { Link } from "react-router-dom";
import { AuthContext } from "../context/AuthContext";
import React, { useEffect, useState } from "react";
import { useSettings } from "../components/pages/Profile-Dependents/Settings/settings-logic/SettingsContext";

const Navbar = () => {
  const { userId } = React.useContext(AuthContext);
  const { settings } = useSettings();
  const [username, setUsername] = useState("");

  useEffect(() => {
    const fetchUsername = async () => {
      if (!userId) return;

      try {
        const response = await fetch(
          `http://localhost:8080/api/users/${userId}`,
          {
            method: "GET",
            headers: {
              "Content-Type": "application/json",
            },
          },
        );

        if (!response.ok) {
          console.error("Failed to fetch username");
          return;
        }

        const data = await response.json();
        setUsername(data.email);
      } catch (err) {
        console.error(err);
      }
    };

    fetchUsername();
  });
  return (
    <nav
      className={`absolute top-0 left-0 right-0 flex justify-between items-center p-4 max-w-full w-full border-b-2 ${
        settings.useLiquidGlass
          ? "bg-white/30 backdrop-blur-lg border-white/30 shadow-sm shadow-white/20"
          : "bg-linear-to-b from-gray-700 to-gray-500 [border-image:linear-gradient(to_right,var(--color-green-400),var(--color-green-600))_1]"
      }`}
    >
      <div className="navbar-left"></div>
      <div className="navbar-center">
        <ul
          className={`nav-links list-none flex m-0 p-0 gap-20 ${settings.useLiquidGlass ? "rounded-lg bg-white/30 backdrop-blur-lg border-white/30 shadow-lg shadow-white/20" : ""}`}
        >
          <li
            className={`m-4 ${settings.useLiquidGlass ? "p-1 hover:backdrop-blur-lg hover:rounded-sm hover:shadow-sm hover:shadow-white/20 hover:text-gray-700" : "hover:text-green-400 "} transition-colors duration-300`}
          >
            <Link to="/app/about">
              <Label>About Us</Label>
            </Link>
          </li>
          <li
            className={`m-4 p-1 ${settings.useLiquidGlass ? " hover:backdrop-blur-lg hover:rounded-sm hover:shadow-sm hover:shadow-white/20 hover:text-gray-700" : "hover:text-green-400 "} transition-colors duration-300`}
          >
            <Link to="/app/gallery">
              <Label>Gallery</Label>
            </Link>
          </li>
          <li
            className={`m-4 p-1  ${settings.useLiquidGlass ? "hover:backdrop-blur-lg hover:rounded-sm hover:shadow-sm hover:shadow-white/20 hover:text-gray-700" : "hover:text-green-400 "} transition-colors duration-300`}
          >
            <Link to="/app/releases">
              <Label>Releases</Label>
            </Link>
          </li>
          <li
            className={`m-4 p-1 ${settings.useLiquidGlass ? " hover:backdrop-blur-lg hover:rounded-sm hover:shadow-sm hover:shadow-white/20 hover:text-gray-700" : "hover:text-green-400 "} transition-colors duration-300`}
          >
            <Link to="/app/webstore">
              <Label>Webstore</Label>
            </Link>
          </li>
          <li
            className={`m-4 p-1 ${settings.useLiquidGlass ? " hover:backdrop-blur-lg hover:rounded-sm hover:shadow-sm hover:shadow-white/20 hover:text-gray-700" : "hover:text-green-400 "} transition-colors duration-300`}
          >
            <Link to="/app/news">
              <Label>News</Label>
            </Link>
          </li>
        </ul>
      </div>
      <div className="flex items-center gap-4">
        <Label>Logged in as {username}</Label>
        <AccountMenu />
      </div>
    </nav>
  );
};

export default Navbar;
