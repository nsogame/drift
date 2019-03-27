package ddl

type BaseGenerator struct {
}

func (g *BaseGenerator) VisitCreateTable(name string) string {
	return ""
}

func (g *BaseGenerator) VisitCreateTableIfNotExists(name string) string {
	return ""
}

func (g *BaseGenerator) VisitDropTable(name string) string {
	return ""
}

func (g *BaseGenerator) VisitDropTableIfExists(name string) string {
	return ""
}

func (g *BaseGenerator) VisitRenameTable(name string) string {
	return ""
}

func (g *BaseGenerator) VisitAlterTable(name string) string {
	return ""
}

func (g *BaseGenerator) VisitAddColumn(name string) string {
	return ""
}

func (g *BaseGenerator) VisitDropColumn(name string) string {
	return ""
}

func (g *BaseGenerator) VisitRenameColumn(name string) string {
	return ""
}
