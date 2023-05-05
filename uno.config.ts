import {
  defineConfig,
  presetTagify,
  presetIcons,
  presetAttributify,
  presetWebFonts,
  presetUno,
  transformerVariantGroup,
  transformerDirectives,
} from "unocss";
import extractorSvelte from "@unocss/extractor-svelte";

export default defineConfig({
  extractors: [extractorSvelte()],
  shortcuts: {
    "flex-centered": "justify-center items-center",
    "size-screen": "w-full min-h-screen",
    card: "bg-white rounded-lg shadow",
    btn: "rounded p-2 all:duration-200 duration-200 will-change-a",
    "btn-primary":
      "bg-dark-400 hover:bg-dark-200 dark:bg-white hover:dark:bg-light-900 text-white dark:text-dark-900",
  },
  presets: [
    presetUno(),
    presetIcons(),
    presetAttributify({
      prefixedOnly: true,
      prefix: "data-",
    }),
    presetTagify({
      prefix: "un-",
    }),
    presetWebFonts({
      fonts: {
        sans: "Source Sans Pro:300,400,500,700",
        mono: ["Fira Code", "Fira Mono:400,700"],
      },
    }),
  ],
  transformers: [transformerVariantGroup(), transformerDirectives()],
});
