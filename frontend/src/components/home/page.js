import style from './style.css';

const Page = () => {
  function getCurrentPageNum() {
    const url = new URL(window.location.href);
    const searchParams = url.searchParams;
    const page = searchParams.get('page');
    return page ? parseInt(page) : 1;
  }

  function getPrevPage() {
    const url = new URL(window.location.href);
    const searchParams = url.searchParams;
    const currentPage = getCurrentPageNum();
    if (currentPage === 1) {
      return null;
    }

    searchParams.set('page', currentPage - 1);
    return `${url.pathname}?${searchParams.toString()}`;
  }

  function getNextPage() {
    const url = new URL(window.location.href);
    const searchParams = url.searchParams;
    const currentPage = getCurrentPageNum();
    searchParams.set('page', currentPage + 1);
    return `${url.pathname}?${searchParams.toString()}`;
  }

  return (
    <div class={style.paging}>
      {getPrevPage() ? <a href={getPrevPage()}>&larr;前ページ</a> : <div />}
      {getNextPage() ? <a href={getNextPage()}>次ページ&rarr;</a> : <div />}
    </div>
  );
};

export default Page;
