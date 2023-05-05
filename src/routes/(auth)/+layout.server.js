import { redirect } from "@sveltejs/kit";

export function load({ locals }) {
  if (locals.user != null) {
    throw redirect(303, "/");
  }
}
