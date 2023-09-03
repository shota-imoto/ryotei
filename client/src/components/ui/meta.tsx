import Head from 'next/head';

type Props = {
  title?: string;
};

/**
 * 追加したいメタ情報があれば随時更新
 */
export const Meta = ({ title }: Props) => {
  return <Head>{title && <title>{title}</title>}</Head>;
};