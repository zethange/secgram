import { SHA1, enc } from 'crypto-js'

const generateGradient = (input: string): string => {
  const hash = SHA1(input).toString(enc.Hex);

  const color1 = `#${hash.slice(0, 6)}`;
  const color2 = `#${hash.slice(6, 12)}`;

  return `linear-gradient(to bottom right, ${color1}, ${color2})`;
}

export { generateGradient }