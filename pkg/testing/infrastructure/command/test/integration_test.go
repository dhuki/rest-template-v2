package test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/dhuki/rest-template-v2/common"
	"github.com/dhuki/rest-template-v2/config"
	"github.com/dhuki/rest-template-v2/pkg/testing/domain/entity"
	"github.com/dhuki/rest-template-v2/pkg/testing/infrastructure/command"
)

// this is testing with real connection database

// command to get coverprofile if test in different package
// go test -coverpkg ./... ./test -coverprofile ./test/coverage.out
// to show in html format
// go tool cover -html=cp.out

var testTableRepo command.TestTableCommand

func TestMain(m *testing.M) {
	err := common.LoadCons(common.ENV_PATH_TESTING)
	if err != nil {
		fmt.Println(err)
		return
	}

	db, err := config.NewDatabase()
	if err != nil {
		fmt.Println(err)
		return
	}

	testTableRepo = command.NewTestTableCommand(db)

	os.Exit(m.Run())
}

func TestIntegrateGetAll(t *testing.T) {
	_, err := testTableRepo.GetAll(context.TODO())
	if err != nil {
		t.Errorf("%s", err)
		return
	}
}

func TestIntegrateCreate(t *testing.T) {
	actual := []entity.TestTable{
		{
			Name:        "test",
			Description: "test",
		},
		{
			Name: "teting with long text",
			Description: `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi accumsan nisi faucibus, mollis enim sed, tincidunt tellus. Nunc vel vehicula nisl. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. In malesuada, leo in fermentum lobortis, dui ipsum imperdiet tellus, ut viverra tellus tortor vel purus. Aenean massa risus, feugiat a eros sit amet, sodales tincidunt sapien. Morbi tincidunt a mauris ac semper. Sed leo magna, facilisis placerat sem ut, ultrices eleifend libero. Sed id pretium neque. Nunc non nibh eget mauris commodo dignissim. Proin eget est eget ligula feugiat cursus nec nec erat. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris non leo id lacus vehicula efficitur. Suspendisse vel enim eget ligula accumsan scelerisque. Quisque elementum nibh tortor, vel lacinia nunc dignissim eu. Etiam ac enim arcu. Sed vel mattis nisl, a mollis nunc.
			Nam eget semper tortor. Nullam posuere purus in urna molestie tincidunt. Sed elementum finibus dolor ut posuere. Sed quis consectetur tortor. Vestibulum posuere placerat sem. Etiam suscipit, ligula et sodales efficitur, arcu massa ultricies orci, euismod eleifend lectus risus ut felis. Pellentesque interdum erat nec risus pellentesque tincidunt. Ut tortor quam, laoreet a massa nec, interdum pharetra augue. Maecenas tincidunt lacus eu quam imperdiet rutrum. Proin sit amet luctus libero. Phasellus non iaculis ligula. Quisque iaculis purus lorem, non auctor erat efficitur sed. Aenean id justo vehicula, interdum dolor at, posuere risus. Nam elementum ante nisl, quis hendrerit sapien finibus ac. Fusce sagittis, quam et finibus euismod, velit felis vehicula nisi, posuere luctus turpis urna pretium arcu. In hac habitasse platea dictumst.
			Praesent egestas auctor mollis. Duis rhoncus cursus libero ut ultrices. In hac habitasse platea dictumst. Nam nibh urna, facilisis id lacinia eget, venenatis eget eros. Mauris condimentum maximus magna, sed vehicula nisl fermentum sed. Aenean dolor arcu, blandit sit amet magna vitae, varius euismod dolor. Donec imperdiet turpis nec sem lobortis euismod. Etiam facilisis diam a pretium dapibus. Phasellus ut ligula blandit ligula varius malesuada et quis metus. Donec placerat sodales eleifend. Integer viverra orci nisl. Etiam nisl elit, iaculis nec mi quis, euismod feugiat metus. Vivamus maximus, augue a vehicula consequat, nulla elit tempus lorem, et rhoncus elit ex ut metus.
			Pellentesque mollis tempor erat eget fringilla. Aenean eu nunc ac metus pellentesque ultrices eu ac dui. Nam purus orci, bibendum non mattis et, lacinia quis nibh. Etiam purus nibh, viverra quis enim ut, malesuada elementum augue. Quisque sodales consectetur viverra. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Etiam ac turpis nulla. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Proin eu tincidunt purus. Sed vel ex enim.
			Curabitur sollicitudin nisl eu diam scelerisque, id gravida nunc fringilla. Quisque lorem tortor, cursus non maximus et, dapibus quis libero. Nunc ipsum nulla, blandit et eros vitae, ultrices efficitur erat. In facilisis diam quam, in gravida leo gravida at. Fusce aliquet urna in feugiat varius. Fusce augue turpis, iaculis nec rutrum quis, imperdiet ac diam. Quisque et dapibus nisl. Etiam quam libero, sagittis sit amet velit vitae, vehicula consequat diam. Integer eu risus libero. Nullam ac dolor ut dui dapibus consequat sodales id leo. Sed ornare dictum risus quis elementum. Maecenas eu tempor ipsum.`,
		},
	}

	for i := range actual {
		err := testTableRepo.Create(context.TODO(), actual[i])
		if err != nil {
			t.Errorf("%s", err)
		}
	}
}
