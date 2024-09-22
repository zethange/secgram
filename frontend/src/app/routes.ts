import { RouteDefinition } from "@solidjs/router";
import { HomeView } from '@/pages/HomeView';

export const routes = [{
  path: "/",
  component: HomeView,
}] satisfies RouteDefinition[]