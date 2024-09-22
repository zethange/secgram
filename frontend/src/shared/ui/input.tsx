import { Component, JSX } from 'solid-js';

export const Input: Component<JSX.InputHTMLAttributes<HTMLInputElement>> = (props) => {
  return <input {...props} class={'border border-gray-300 rounded-md shadow-sm focus:ring focus:ring-blue-500 focus:border-blue-500 p-2 w-full ' + props.class} />
}