const lightCodeTheme = require('prism-react-renderer/themes/github');
const darkCodeTheme = require('prism-react-renderer/themes/dracula');

/** @type {import('@docusaurus/types').DocusaurusConfig} */
module.exports = {
  title: 'logstore2parquet',
  tagline: 'Convert logstore_standard_log CSVs to Parquet',
  url: 'https://open-lms-open-source.github.io',
  baseUrl: '/logstore2parquet/',
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',
  favicon: 'img/favicon.ico',
  organizationName: 'open-lms-open-source',
  projectName: 'logstore2parquet',
  themeConfig: {
    navbar: {
      title: 'logstore2parquet',
      logo: {
        src: 'img/logo.svg',
      },
      items: [
        {
          type: 'doc',
          docId: 'introduction',
          position: 'right',
          label: 'Docs',
        },
        {
          to: 'blog',
          position: 'right',
          label: 'Release Notes',
        },
        {
          href: 'https://github.com/open-lms-open-source/logstore2parquet',
          className: 'header-github-link',
          position: 'right',
          title: 'View on GitHub',
        },
      ],
    },
    footer: {
      style: 'dark',
      links: [
      ],
      copyright: `Copyright © ${new Date().getFullYear()} Open LMS.<br>Built with Docusaurus.<br>Icons made by <a href="https://www.freepik.com" title="Freepik">Freepik</a> from <a href="https://www.flaticon.com/" title="Flaticon">www.flaticon.com</a>.`,
    },
    prism: {
      theme: lightCodeTheme,
      darkTheme: darkCodeTheme,
    },
  },
  presets: [
    [
      '@docusaurus/preset-classic',
      {
        docs: {
          sidebarPath: require.resolve('./sidebars.js'),
          editUrl:
            'https://github.com/open-lms-open-source/logstore2parquet/edit/main/website/',
        },
        blog: {
          showReadingTime: true,
          editUrl:
            'https://github.com/open-lms-open-source/logstore2parquet/edit/main/website/release-notes/',
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      },
    ],
  ],
};
