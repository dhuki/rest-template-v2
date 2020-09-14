/*
 Navicat Premium Data Transfer

 Source Server         : belajar
 Source Server Type    : PostgreSQL
 Source Server Version : 110001
 Source Host           : localhost:5432
 Source Catalog        : testing
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 110001
 File Encoding         : 65001

 Date: 01/09/2020 16:47:21

 And don't forget to create sequence for primary key to perform increment automatically
 with name : 'test_table_id_seq'
*/


-- ----------------------------
-- Table structure for test_tables
-- ----------------------------
DROP TABLE IF EXISTS "public"."test_tables";
CREATE TABLE "public"."test_tables" (
  "id" int4 NOT NULL DEFAULT nextval('test_table_id_seq'::regclass),
  "name" varchar(100) COLLATE "pg_catalog"."default",
  "description" text COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "public"."test_tables"."id" IS ' ';

-- ----------------------------
-- Records of test_tables
-- ----------------------------
INSERT INTO "public"."test_tables" VALUES (1, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (2, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (3, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (4, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (5, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (6, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (7, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (8, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (9, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (10, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (11, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (12, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (13, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (14, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (15, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (16, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (17, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (18, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (19, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (20, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (21, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (22, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (23, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (24, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (25, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (26, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (27, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (28, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (29, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (30, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (31, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (32, 'test', 'test');
INSERT INTO "public"."test_tables" VALUES (33, 'test', 'test');
INSERT INTO "public"."test_tables" VALUES (34, 'test', 'test');
INSERT INTO "public"."test_tables" VALUES (35, 'test', 'test');
INSERT INTO "public"."test_tables" VALUES (36, 'test', 'test');
INSERT INTO "public"."test_tables" VALUES (37, 'test', 'test');
INSERT INTO "public"."test_tables" VALUES (38, 'test', 'test');
INSERT INTO "public"."test_tables" VALUES (39, 'teting with long text', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi accumsan nisi faucibus, mollis enim sed, tincidunt tellus. Nunc vel vehicula nisl. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. In malesuada, leo in fermentum lobortis, dui ipsum imperdiet tellus, ut viverra tellus tortor vel purus. Aenean massa risus, feugiat a eros sit amet, sodales tincidunt sapien. Morbi tincidunt a mauris ac semper. Sed leo magna, facilisis placerat sem ut, ultrices eleifend libero. Sed id pretium neque. Nunc non nibh eget mauris commodo dignissim. Proin eget est eget ligula feugiat cursus nec nec erat. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris non leo id lacus vehicula efficitur. Suspendisse vel enim eget ligula accumsan scelerisque. Quisque elementum nibh tortor, vel lacinia nunc dignissim eu. Etiam ac enim arcu. Sed vel mattis nisl, a mollis nunc.
			Nam eget semper tortor. Nullam posuere purus in urna molestie tincidunt. Sed elementum finibus dolor ut posuere. Sed quis consectetur tortor. Vestibulum posuere placerat sem. Etiam suscipit, ligula et sodales efficitur, arcu massa ultricies orci, euismod eleifend lectus risus ut felis. Pellentesque interdum erat nec risus pellentesque tincidunt. Ut tortor quam, laoreet a massa nec, interdum pharetra augue. Maecenas tincidunt lacus eu quam imperdiet rutrum. Proin sit amet luctus libero. Phasellus non iaculis ligula. Quisque iaculis purus lorem, non auctor erat efficitur sed. Aenean id justo vehicula, interdum dolor at, posuere risus. Nam elementum ante nisl, quis hendrerit sapien finibus ac. Fusce sagittis, quam et finibus euismod, velit felis vehicula nisi, posuere luctus turpis urna pretium arcu. In hac habitasse platea dictumst.
			Praesent egestas auctor mollis. Duis rhoncus cursus libero ut ultrices. In hac habitasse platea dictumst. Nam nibh urna, facilisis id lacinia eget, venenatis eget eros. Mauris condimentum maximus magna, sed vehicula nisl fermentum sed. Aenean dolor arcu, blandit sit amet magna vitae, varius euismod dolor. Donec imperdiet turpis nec sem lobortis euismod. Etiam facilisis diam a pretium dapibus. Phasellus ut ligula blandit ligula varius malesuada et quis metus. Donec placerat sodales eleifend. Integer viverra orci nisl. Etiam nisl elit, iaculis nec mi quis, euismod feugiat metus. Vivamus maximus, augue a vehicula consequat, nulla elit tempus lorem, et rhoncus elit ex ut metus.
			Pellentesque mollis tempor erat eget fringilla. Aenean eu nunc ac metus pellentesque ultrices eu ac dui. Nam purus orci, bibendum non mattis et, lacinia quis nibh. Etiam purus nibh, viverra quis enim ut, malesuada elementum augue. Quisque sodales consectetur viverra. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Etiam ac turpis nulla. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Proin eu tincidunt purus. Sed vel ex enim.
			Curabitur sollicitudin nisl eu diam scelerisque, id gravida nunc fringilla. Quisque lorem tortor, cursus non maximus et, dapibus quis libero. Nunc ipsum nulla, blandit et eros vitae, ultrices efficitur erat. In facilisis diam quam, in gravida leo gravida at. Fusce aliquet urna in feugiat varius. Fusce augue turpis, iaculis nec rutrum quis, imperdiet ac diam. Quisque et dapibus nisl. Etiam quam libero, sagittis sit amet velit vitae, vehicula consequat diam. Integer eu risus libero. Nullam ac dolor ut dui dapibus consequat sodales id leo. Sed ornare dictum risus quis elementum. Maecenas eu tempor ipsum.');
INSERT INTO "public"."test_tables" VALUES (40, 'test', 'test');
INSERT INTO "public"."test_tables" VALUES (41, 'teting with long text', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi accumsan nisi faucibus, mollis enim sed, tincidunt tellus. Nunc vel vehicula nisl. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. In malesuada, leo in fermentum lobortis, dui ipsum imperdiet tellus, ut viverra tellus tortor vel purus. Aenean massa risus, feugiat a eros sit amet, sodales tincidunt sapien. Morbi tincidunt a mauris ac semper. Sed leo magna, facilisis placerat sem ut, ultrices eleifend libero. Sed id pretium neque. Nunc non nibh eget mauris commodo dignissim. Proin eget est eget ligula feugiat cursus nec nec erat. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris non leo id lacus vehicula efficitur. Suspendisse vel enim eget ligula accumsan scelerisque. Quisque elementum nibh tortor, vel lacinia nunc dignissim eu. Etiam ac enim arcu. Sed vel mattis nisl, a mollis nunc.
			Nam eget semper tortor. Nullam posuere purus in urna molestie tincidunt. Sed elementum finibus dolor ut posuere. Sed quis consectetur tortor. Vestibulum posuere placerat sem. Etiam suscipit, ligula et sodales efficitur, arcu massa ultricies orci, euismod eleifend lectus risus ut felis. Pellentesque interdum erat nec risus pellentesque tincidunt. Ut tortor quam, laoreet a massa nec, interdum pharetra augue. Maecenas tincidunt lacus eu quam imperdiet rutrum. Proin sit amet luctus libero. Phasellus non iaculis ligula. Quisque iaculis purus lorem, non auctor erat efficitur sed. Aenean id justo vehicula, interdum dolor at, posuere risus. Nam elementum ante nisl, quis hendrerit sapien finibus ac. Fusce sagittis, quam et finibus euismod, velit felis vehicula nisi, posuere luctus turpis urna pretium arcu. In hac habitasse platea dictumst.
			Praesent egestas auctor mollis. Duis rhoncus cursus libero ut ultrices. In hac habitasse platea dictumst. Nam nibh urna, facilisis id lacinia eget, venenatis eget eros. Mauris condimentum maximus magna, sed vehicula nisl fermentum sed. Aenean dolor arcu, blandit sit amet magna vitae, varius euismod dolor. Donec imperdiet turpis nec sem lobortis euismod. Etiam facilisis diam a pretium dapibus. Phasellus ut ligula blandit ligula varius malesuada et quis metus. Donec placerat sodales eleifend. Integer viverra orci nisl. Etiam nisl elit, iaculis nec mi quis, euismod feugiat metus. Vivamus maximus, augue a vehicula consequat, nulla elit tempus lorem, et rhoncus elit ex ut metus.
			Pellentesque mollis tempor erat eget fringilla. Aenean eu nunc ac metus pellentesque ultrices eu ac dui. Nam purus orci, bibendum non mattis et, lacinia quis nibh. Etiam purus nibh, viverra quis enim ut, malesuada elementum augue. Quisque sodales consectetur viverra. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Etiam ac turpis nulla. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Proin eu tincidunt purus. Sed vel ex enim.
			Curabitur sollicitudin nisl eu diam scelerisque, id gravida nunc fringilla. Quisque lorem tortor, cursus non maximus et, dapibus quis libero. Nunc ipsum nulla, blandit et eros vitae, ultrices efficitur erat. In facilisis diam quam, in gravida leo gravida at. Fusce aliquet urna in feugiat varius. Fusce augue turpis, iaculis nec rutrum quis, imperdiet ac diam. Quisque et dapibus nisl. Etiam quam libero, sagittis sit amet velit vitae, vehicula consequat diam. Integer eu risus libero. Nullam ac dolor ut dui dapibus consequat sodales id leo. Sed ornare dictum risus quis elementum. Maecenas eu tempor ipsum.');
INSERT INTO "public"."test_tables" VALUES (42, 'test', 'test');
INSERT INTO "public"."test_tables" VALUES (89, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (90, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (91, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (92, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (93, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (43, 'teting with long text', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi accumsan nisi faucibus, mollis enim sed, tincidunt tellus. Nunc vel vehicula nisl. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. In malesuada, leo in fermentum lobortis, dui ipsum imperdiet tellus, ut viverra tellus tortor vel purus. Aenean massa risus, feugiat a eros sit amet, sodales tincidunt sapien. Morbi tincidunt a mauris ac semper. Sed leo magna, facilisis placerat sem ut, ultrices eleifend libero. Sed id pretium neque. Nunc non nibh eget mauris commodo dignissim. Proin eget est eget ligula feugiat cursus nec nec erat. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris non leo id lacus vehicula efficitur. Suspendisse vel enim eget ligula accumsan scelerisque. Quisque elementum nibh tortor, vel lacinia nunc dignissim eu. Etiam ac enim arcu. Sed vel mattis nisl, a mollis nunc.
			Nam eget semper tortor. Nullam posuere purus in urna molestie tincidunt. Sed elementum finibus dolor ut posuere. Sed quis consectetur tortor. Vestibulum posuere placerat sem. Etiam suscipit, ligula et sodales efficitur, arcu massa ultricies orci, euismod eleifend lectus risus ut felis. Pellentesque interdum erat nec risus pellentesque tincidunt. Ut tortor quam, laoreet a massa nec, interdum pharetra augue. Maecenas tincidunt lacus eu quam imperdiet rutrum. Proin sit amet luctus libero. Phasellus non iaculis ligula. Quisque iaculis purus lorem, non auctor erat efficitur sed. Aenean id justo vehicula, interdum dolor at, posuere risus. Nam elementum ante nisl, quis hendrerit sapien finibus ac. Fusce sagittis, quam et finibus euismod, velit felis vehicula nisi, posuere luctus turpis urna pretium arcu. In hac habitasse platea dictumst.
			Praesent egestas auctor mollis. Duis rhoncus cursus libero ut ultrices. In hac habitasse platea dictumst. Nam nibh urna, facilisis id lacinia eget, venenatis eget eros. Mauris condimentum maximus magna, sed vehicula nisl fermentum sed. Aenean dolor arcu, blandit sit amet magna vitae, varius euismod dolor. Donec imperdiet turpis nec sem lobortis euismod. Etiam facilisis diam a pretium dapibus. Phasellus ut ligula blandit ligula varius malesuada et quis metus. Donec placerat sodales eleifend. Integer viverra orci nisl. Etiam nisl elit, iaculis nec mi quis, euismod feugiat metus. Vivamus maximus, augue a vehicula consequat, nulla elit tempus lorem, et rhoncus elit ex ut metus.
			Pellentesque mollis tempor erat eget fringilla. Aenean eu nunc ac metus pellentesque ultrices eu ac dui. Nam purus orci, bibendum non mattis et, lacinia quis nibh. Etiam purus nibh, viverra quis enim ut, malesuada elementum augue. Quisque sodales consectetur viverra. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Etiam ac turpis nulla. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Proin eu tincidunt purus. Sed vel ex enim.
			Curabitur sollicitudin nisl eu diam scelerisque, id gravida nunc fringilla. Quisque lorem tortor, cursus non maximus et, dapibus quis libero. Nunc ipsum nulla, blandit et eros vitae, ultrices efficitur erat. In facilisis diam quam, in gravida leo gravida at. Fusce aliquet urna in feugiat varius. Fusce augue turpis, iaculis nec rutrum quis, imperdiet ac diam. Quisque et dapibus nisl. Etiam quam libero, sagittis sit amet velit vitae, vehicula consequat diam. Integer eu risus libero. Nullam ac dolor ut dui dapibus consequat sodales id leo. Sed ornare dictum risus quis elementum. Maecenas eu tempor ipsum.');
INSERT INTO "public"."test_tables" VALUES (44, 'test', 'test');
INSERT INTO "public"."test_tables" VALUES (45, 'teting with long text', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi accumsan nisi faucibus, mollis enim sed, tincidunt tellus. Nunc vel vehicula nisl. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. In malesuada, leo in fermentum lobortis, dui ipsum imperdiet tellus, ut viverra tellus tortor vel purus. Aenean massa risus, feugiat a eros sit amet, sodales tincidunt sapien. Morbi tincidunt a mauris ac semper. Sed leo magna, facilisis placerat sem ut, ultrices eleifend libero. Sed id pretium neque. Nunc non nibh eget mauris commodo dignissim. Proin eget est eget ligula feugiat cursus nec nec erat. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris non leo id lacus vehicula efficitur. Suspendisse vel enim eget ligula accumsan scelerisque. Quisque elementum nibh tortor, vel lacinia nunc dignissim eu. Etiam ac enim arcu. Sed vel mattis nisl, a mollis nunc.
			Nam eget semper tortor. Nullam posuere purus in urna molestie tincidunt. Sed elementum finibus dolor ut posuere. Sed quis consectetur tortor. Vestibulum posuere placerat sem. Etiam suscipit, ligula et sodales efficitur, arcu massa ultricies orci, euismod eleifend lectus risus ut felis. Pellentesque interdum erat nec risus pellentesque tincidunt. Ut tortor quam, laoreet a massa nec, interdum pharetra augue. Maecenas tincidunt lacus eu quam imperdiet rutrum. Proin sit amet luctus libero. Phasellus non iaculis ligula. Quisque iaculis purus lorem, non auctor erat efficitur sed. Aenean id justo vehicula, interdum dolor at, posuere risus. Nam elementum ante nisl, quis hendrerit sapien finibus ac. Fusce sagittis, quam et finibus euismod, velit felis vehicula nisi, posuere luctus turpis urna pretium arcu. In hac habitasse platea dictumst.
			Praesent egestas auctor mollis. Duis rhoncus cursus libero ut ultrices. In hac habitasse platea dictumst. Nam nibh urna, facilisis id lacinia eget, venenatis eget eros. Mauris condimentum maximus magna, sed vehicula nisl fermentum sed. Aenean dolor arcu, blandit sit amet magna vitae, varius euismod dolor. Donec imperdiet turpis nec sem lobortis euismod. Etiam facilisis diam a pretium dapibus. Phasellus ut ligula blandit ligula varius malesuada et quis metus. Donec placerat sodales eleifend. Integer viverra orci nisl. Etiam nisl elit, iaculis nec mi quis, euismod feugiat metus. Vivamus maximus, augue a vehicula consequat, nulla elit tempus lorem, et rhoncus elit ex ut metus.
			Pellentesque mollis tempor erat eget fringilla. Aenean eu nunc ac metus pellentesque ultrices eu ac dui. Nam purus orci, bibendum non mattis et, lacinia quis nibh. Etiam purus nibh, viverra quis enim ut, malesuada elementum augue. Quisque sodales consectetur viverra. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Etiam ac turpis nulla. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Proin eu tincidunt purus. Sed vel ex enim.
			Curabitur sollicitudin nisl eu diam scelerisque, id gravida nunc fringilla. Quisque lorem tortor, cursus non maximus et, dapibus quis libero. Nunc ipsum nulla, blandit et eros vitae, ultrices efficitur erat. In facilisis diam quam, in gravida leo gravida at. Fusce aliquet urna in feugiat varius. Fusce augue turpis, iaculis nec rutrum quis, imperdiet ac diam. Quisque et dapibus nisl. Etiam quam libero, sagittis sit amet velit vitae, vehicula consequat diam. Integer eu risus libero. Nullam ac dolor ut dui dapibus consequat sodales id leo. Sed ornare dictum risus quis elementum. Maecenas eu tempor ipsum.');
INSERT INTO "public"."test_tables" VALUES (46, 'test', 'test');
INSERT INTO "public"."test_tables" VALUES (47, 'teting with long text', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi accumsan nisi faucibus, mollis enim sed, tincidunt tellus. Nunc vel vehicula nisl. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. In malesuada, leo in fermentum lobortis, dui ipsum imperdiet tellus, ut viverra tellus tortor vel purus. Aenean massa risus, feugiat a eros sit amet, sodales tincidunt sapien. Morbi tincidunt a mauris ac semper. Sed leo magna, facilisis placerat sem ut, ultrices eleifend libero. Sed id pretium neque. Nunc non nibh eget mauris commodo dignissim. Proin eget est eget ligula feugiat cursus nec nec erat. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris non leo id lacus vehicula efficitur. Suspendisse vel enim eget ligula accumsan scelerisque. Quisque elementum nibh tortor, vel lacinia nunc dignissim eu. Etiam ac enim arcu. Sed vel mattis nisl, a mollis nunc.
			Nam eget semper tortor. Nullam posuere purus in urna molestie tincidunt. Sed elementum finibus dolor ut posuere. Sed quis consectetur tortor. Vestibulum posuere placerat sem. Etiam suscipit, ligula et sodales efficitur, arcu massa ultricies orci, euismod eleifend lectus risus ut felis. Pellentesque interdum erat nec risus pellentesque tincidunt. Ut tortor quam, laoreet a massa nec, interdum pharetra augue. Maecenas tincidunt lacus eu quam imperdiet rutrum. Proin sit amet luctus libero. Phasellus non iaculis ligula. Quisque iaculis purus lorem, non auctor erat efficitur sed. Aenean id justo vehicula, interdum dolor at, posuere risus. Nam elementum ante nisl, quis hendrerit sapien finibus ac. Fusce sagittis, quam et finibus euismod, velit felis vehicula nisi, posuere luctus turpis urna pretium arcu. In hac habitasse platea dictumst.
			Praesent egestas auctor mollis. Duis rhoncus cursus libero ut ultrices. In hac habitasse platea dictumst. Nam nibh urna, facilisis id lacinia eget, venenatis eget eros. Mauris condimentum maximus magna, sed vehicula nisl fermentum sed. Aenean dolor arcu, blandit sit amet magna vitae, varius euismod dolor. Donec imperdiet turpis nec sem lobortis euismod. Etiam facilisis diam a pretium dapibus. Phasellus ut ligula blandit ligula varius malesuada et quis metus. Donec placerat sodales eleifend. Integer viverra orci nisl. Etiam nisl elit, iaculis nec mi quis, euismod feugiat metus. Vivamus maximus, augue a vehicula consequat, nulla elit tempus lorem, et rhoncus elit ex ut metus.
			Pellentesque mollis tempor erat eget fringilla. Aenean eu nunc ac metus pellentesque ultrices eu ac dui. Nam purus orci, bibendum non mattis et, lacinia quis nibh. Etiam purus nibh, viverra quis enim ut, malesuada elementum augue. Quisque sodales consectetur viverra. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Etiam ac turpis nulla. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Proin eu tincidunt purus. Sed vel ex enim.
			Curabitur sollicitudin nisl eu diam scelerisque, id gravida nunc fringilla. Quisque lorem tortor, cursus non maximus et, dapibus quis libero. Nunc ipsum nulla, blandit et eros vitae, ultrices efficitur erat. In facilisis diam quam, in gravida leo gravida at. Fusce aliquet urna in feugiat varius. Fusce augue turpis, iaculis nec rutrum quis, imperdiet ac diam. Quisque et dapibus nisl. Etiam quam libero, sagittis sit amet velit vitae, vehicula consequat diam. Integer eu risus libero. Nullam ac dolor ut dui dapibus consequat sodales id leo. Sed ornare dictum risus quis elementum. Maecenas eu tempor ipsum.');
INSERT INTO "public"."test_tables" VALUES (48, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (49, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (50, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (51, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (52, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (53, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (54, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (55, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (56, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (57, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (58, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (59, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (60, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (61, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (62, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (63, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (64, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (65, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (66, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (67, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (68, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (69, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (70, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (71, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (72, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (73, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (74, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (75, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (76, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (77, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (78, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (79, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (80, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (81, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (82, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (83, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (84, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (85, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (86, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (87, 'testing', 'testing');
INSERT INTO "public"."test_tables" VALUES (88, 'testing', 'testing');

-- ----------------------------
-- Primary Key structure for table test_tables
-- ----------------------------
ALTER TABLE "public"."test_tables" ADD CONSTRAINT "test_table_pkey" PRIMARY KEY ("id");
