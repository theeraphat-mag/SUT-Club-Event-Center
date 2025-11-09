import { useRoutes } from "react-router-dom"; // ลบ RouteObject ออกจากตรงนี้
import type { RouteObject } from "react-router-dom"; // Import RouteObject เป็น type
import AdminRoutes from "./AdminRoutes";
import MainRoutes from "./MainRoutes";

function ConfigRoutes() {
  const isLoggedIn = localStorage.getItem("isLogin") === "true";
  let routes: RouteObject[] = []; // ใช้ RouteObject เป็น type ตรงนี้
  if (isLoggedIn) {
    routes = [AdminRoutes(isLoggedIn), MainRoutes()];
  } else {
    routes = [MainRoutes()];
  }
  return useRoutes(routes);
}

export default ConfigRoutes;
