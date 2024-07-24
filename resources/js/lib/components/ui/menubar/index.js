import { Menubar as MenubarPrimitive } from "bits-ui";
import Root from "./menubar.svelte";
import Content from "./menubar-content.svelte";
import Item from "./menubar-item.svelte";
import Separator from "./menubar-separator.svelte";
import SubContent from "./menubar-sub-content.svelte";
import SubTrigger from "./menubar-sub-trigger.svelte";
import Trigger from "./menubar-trigger.svelte";
const Menu = MenubarPrimitive.Menu;
const Sub = MenubarPrimitive.Sub;
export {
  Root,
  Content,
  Item,
  Separator,
  SubContent,
  SubTrigger,
  Trigger,
  Menu,
  Sub,
  //
  Root as Menubar,
  Content as MenubarContent,
  Item as MenubarItem,
  Separator as MenubarSeparator,
  SubContent as MenubarSubContent,
  SubTrigger as MenubarSubTrigger,
  Trigger as MenubarTrigger,
  Menu as MenubarMenu,
  Sub as MenubarSub,
};
