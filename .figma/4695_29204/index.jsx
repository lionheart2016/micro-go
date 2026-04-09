import React from 'react';

import styles from './index.module.scss';

const Component = () => {
  return (
    <div className={styles.dashboard}>
      <div className={styles.container}>
        <p className={styles.text}>概览</p>
        <div className={styles.paragraph}>
          <p className={styles.text2}>XAUa 黄金代币发行管理</p>
        </div>
      </div>
      <div className={styles.container11}>
        <div className={styles.container4}>
          <div className={styles.container3}>
            <div className={styles.text3}>
              <p className={styles.text2}>待审批批次</p>
            </div>
            <div className={styles.container2}>
              <img src="../image/mnnyjjzj-phdratq.svg" className={styles.icon} />
            </div>
          </div>
          <p className={styles.a2}>2</p>
        </div>
        <div className={styles.container7}>
          <div className={styles.container6}>
            <div className={styles.text4}>
              <p className={styles.text2}>今日认购 (USDT)</p>
            </div>
            <div className={styles.container5}>
              <img src="../image/mnnyjjzj-2obdiof.svg" className={styles.icon} />
            </div>
          </div>
          <p className={styles.a2}>285,000</p>
        </div>
        <div className={styles.container4}>
          <div className={styles.container3}>
            <div className={styles.text3}>
              <p className={styles.text2}>累计发行 XAUa</p>
            </div>
            <div className={styles.container2}>
              <img src="../image/mnnyjjzj-4hdwxyb.svg" className={styles.icon} />
            </div>
          </div>
          <p className={styles.a2}>137.22</p>
        </div>
        <div className={styles.container10}>
          <div className={styles.container9}>
            <div className={styles.text4}>
              <p className={styles.text2}>累计赎回 (USDT)</p>
            </div>
            <div className={styles.container8}>
              <img src="../image/mnnyjjzj-2c36qrq.svg" className={styles.icon} />
            </div>
          </div>
          <p className={styles.a2}>221,880</p>
        </div>
      </div>
      <div className={styles.container14}>
        <img src="../image/mnnyjjzj-9rfpmww.svg" className={styles.icon2} />
        <div className={styles.container13}>
          <div className={styles.container12}>
            <p className={styles.text7}>
              <span className={styles.text5}>有&nbsp;</span>
              <span className={styles.text6}>2</span>
              <span className={styles.text5}>&nbsp;个批次订单等待审批</span>
            </p>
          </div>
          <div className={styles.paragraph2}>
            <p className={styles.text8}>
              来自 YunFeng Securities 的订单已汇总，请前往审批中心处理。
            </p>
          </div>
          <div className={styles.button}>
            <p className={styles.text9}>前往审批 →</p>
          </div>
        </div>
      </div>
      <div className={styles.container15}>
        <div className={styles.heading3}>
          <p className={styles.text10}>最近批次</p>
        </div>
        <div className={styles.table}>
          <div className={styles.tableHeader}>
            <div className={styles.tableRow}>
              <p className={styles.text11}>批次编号</p>
              <p className={styles.text12}>方向</p>
              <p className={styles.text13}>交易日</p>
              <p className={styles.text14}>收取</p>
              <p className={styles.text15}>发送</p>
              <p className={styles.text16}>子订单</p>
              <p className={styles.text17}>状态</p>
            </div>
          </div>
          <div className={styles.tableBody}>
            <div className={styles.tableRow2}>
              <p className={styles.bAtchs20260325001}>BATCH-S-20260325-001</p>
              <div className={styles.text19}>
                <img src="../image/mnnyjjzj-tvd74sn.svg" className={styles.icon3} />
                <p className={styles.text18}>认购</p>
              </div>
              <p className={styles.a20260325}>2026-03-25</p>
              <p className={styles.a285000Usdt}>285,000 USDT</p>
              <p className={styles.a10254XaUa}>102.54 XAUa</p>
              <p className={styles.a5}>5</p>
              <div className={styles.tableCell}>
                <div className={styles.statusBadge}>
                  <p className={styles.text20}>待审批</p>
                </div>
              </div>
            </div>
            <div className={styles.tableRow3}>
              <p className={styles.bAtchs20260325001}>BATCH-R-20260325-001</p>
              <div className={styles.text19}>
                <img src="../image/mnnyjjzj-1s0j4j6.svg" className={styles.icon3} />
                <p className={styles.text18}>赎回</p>
              </div>
              <p className={styles.a20260325}>2026-03-25</p>
              <p className={styles.a4520XaUa}>45.20 XAUa</p>
              <p className={styles.a12562282Usdt}>125,622.82 USDT</p>
              <p className={styles.a5}>3</p>
              <div className={styles.tableCell}>
                <div className={styles.statusBadge}>
                  <p className={styles.text20}>待审批</p>
                </div>
              </div>
            </div>
            <div className={styles.tableRow4}>
              <p className={styles.bAtchs20260325001}>BATCH-S-20260324-001</p>
              <div className={styles.text19}>
                <img src="../image/mnnyjjzj-tvd74sn.svg" className={styles.icon3} />
                <p className={styles.text18}>认购</p>
              </div>
              <p className={styles.a20260325}>2026-03-24</p>
              <p className={styles.a285000Usdt}>520,000 USDT</p>
              <p className={styles.a10254XaUa}>187.42 XAUa</p>
              <p className={styles.a5}>3</p>
              <div className={styles.tableCell2}>
                <div className={styles.statusBadge2}>
                  <p className={styles.text21}>已审批</p>
                </div>
              </div>
            </div>
            <div className={styles.tableRow5}>
              <p className={styles.bAtchs20260325001}>BATCH-R-20260324-001</p>
              <div className={styles.text19}>
                <img src="../image/mnnyjjzj-1s0j4j6.svg" className={styles.icon3} />
                <p className={styles.text18}>赎回</p>
              </div>
              <p className={styles.a20260325}>2026-03-24</p>
              <p className={styles.a4520XaUa}>80.00 XAUa</p>
              <p className={styles.a10254XaUa}>221,880 USDT</p>
              <p className={styles.a5}>2</p>
              <div className={styles.tableCell3}>
                <div className={styles.statusBadge3}>
                  <p className={styles.text22}>已结算</p>
                </div>
              </div>
            </div>
            <div className={styles.tableRow6}>
              <p className={styles.bAtchs20260325001}>BATCH-S-20260323-001</p>
              <div className={styles.text19}>
                <img src="../image/mnnyjjzj-tvd74sn.svg" className={styles.icon3} />
                <p className={styles.text18}>认购</p>
              </div>
              <p className={styles.a20260325}>2026-03-23</p>
              <p className={styles.a285000Usdt}>380,000 USDT</p>
              <p className={styles.a10254XaUa}>137.22 XAUa</p>
              <p className={styles.a5}>2</p>
              <div className={styles.tableCell3}>
                <div className={styles.statusBadge3}>
                  <p className={styles.text22}>已结算</p>
                </div>
              </div>
            </div>
            <div className={styles.tableRow7}>
              <p className={styles.bAtchs20260325001}>BATCH-S-20260322-001</p>
              <div className={styles.text23}>
                <img src="../image/mnnyjjzj-tvd74sn.svg" className={styles.icon4} />
                <p className={styles.text18}>认购</p>
              </div>
              <p className={styles.a20260325}>2026-03-22</p>
              <p className={styles.a4520XaUa}>150,000 USDT</p>
              <p className={styles.a10254XaUa}>54.15 XAUa</p>
              <p className={styles.a5}>1</p>
              <div className={styles.tableCell4}>
                <div className={styles.statusBadge4}>
                  <p className={styles.text24}>已拒绝</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Component;
