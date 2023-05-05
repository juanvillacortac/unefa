/// <reference types="unocss/vite" />

// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
import PocketBase from "pocketbase";

declare global {
  type NullablePartial<T> = {
    [P in keyof T]: T[P] | null;
  };

  namespace App {
    // interface Error {}
    interface Locals {
      pb: PocketBase;
      user: import("pocketbase").Record | import("pocketbase").Admin | null;
    }
    // interface PageData {}
    // interface Platform {}
  }
}

export {};
