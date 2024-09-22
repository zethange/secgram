import { Component, JSX } from 'solid-js';

export const Button: Component<JSX.ButtonHTMLAttributes<HTMLButtonElement>> = (props) => {
  return <button {...props} class={'border font-semibold py-2 px-4 rounded-md transition duration-200 hover:scale-95 ' + props.class} />
}